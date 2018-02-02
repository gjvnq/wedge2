package wedge

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gjvnq/go.uuid"
)

type AccountsDBConn struct{}

var Accounts AccountsDBConn

type Account struct {
	ID                uuid.UUID           `json:"id" gorm:"primary_key"`
	BookID            uuid.UUID           `json:"book_id"`
	ParentID          uuid.UUID           `json:"parent_id"`
	Name              string              `json:"name"`
	LocalBalanceIDs   map[uuid.UUID]int64 `json:"local_balance_ids"`
	LocalBalanceCodes map[string]int64    `json:"local_balance_codes"`
	TotalBalanceIDs   map[uuid.UUID]int64 `json:"total_balance_ids,omitempty"`
	TotalBalanceCodes map[string]int64    `json:"total_balance_codes,omitempty"`
	Historic          []BalanceRecord     `json:"historic"`
	Children          []Account           `json:"children"`
}

type BalanceRecord struct {
	Date         LDate               `json:"date"`
	TotalByIDs   map[uuid.UUID]int64 `json:"total_ids"`
	TotalByCodes map[string]int64    `json:"total_codes"`
	DeltaByIDs   map[uuid.UUID]int64 `json:"delta_ids"`
	DeltaByCodes map[string]int64    `json:"delta_codes"`
}

func NewBalanceRecord(date LDate) BalanceRecord {
	br := BalanceRecord{}
	br.Date = date
	br.TotalByIDs = make(map[uuid.UUID]int64)
	br.TotalByCodes = make(map[string]int64)
	br.DeltaByIDs = make(map[uuid.UUID]int64)
	br.DeltaByCodes = make(map[string]int64)
	return br
}

func (br *BalanceRecord) Yesterday(yesterday_ids map[uuid.UUID]int64, yesterday_codes map[string]int64) {
	for key, val := range yesterday_ids {
		br.TotalByIDs[key] = val
	}
	for key, val := range yesterday_codes {
		br.TotalByCodes[key] = val
	}
}

func (br *BalanceRecord) Add(id uuid.UUID, code string, amount int64) {
	br.DeltaByIDs[id] += amount
	br.DeltaByCodes[code] += amount
	br.TotalByIDs[id] += amount
	br.TotalByCodes[code] += amount
}

func (acc Account) String() string {
	return fmt.Sprintf("<%s %s %d>", acc.ID, acc.Name, len(acc.Children))
}

func (acc *Account) Validate() error {
	acc.Name = strings.TrimSpace(acc.Name)
	if len(acc.Name) == 0 {
		return errors.New("account name must not be empty")
	}
	return nil
}

func (acc *Account) LoadHistoric(from, to LDate) error {
	acc.Historic = make([]BalanceRecord, 0)
	var last_br *BalanceRecord
	err := acc.LoadBalanceAt(from)
	if err != nil {
		return err
	}

	acc.Historic = append(acc.Historic, NewBalanceRecord(from))
	last_br = &acc.Historic[0]
	last_br.Yesterday(acc.LocalBalanceIDs, acc.LocalBalanceCodes)

	yesterday := LDate{}
	// Load deltas
	rows, err := DB.Query("SELECT `AssetID`, `AssetCode`, `MovementDate`, SUM(`Amount`) FROM `movements_view` WHERE `AccountID` = ? AND ? < `MovementDate` AND `MovementDate` <= ? AND `MovementStatus` != 'C' GROUP BY `AccountID`, `AssetID`, `MovementDate` ORDER BY `MovementDate` ASC", acc.ID, from, to)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		Log.WarningF("Error when loading historic: %#v", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var asset_id uuid.UUID
		var asset_code string
		var amount int64
		var mov_date LDate
		err = rows.Scan(&asset_id, &asset_code, &mov_date, &amount)
		if err != nil {
			Log.WarningF("Error when loading historic for account %s: %#v", acc.ID.String(), err)
			return err
		}
		// Do not include zeros
		if amount == 0 {
			continue
		}
		// Add record
		if !mov_date.Equals(yesterday) {
			acc.Historic = append(acc.Historic, NewBalanceRecord(mov_date))
			prev_br := last_br
			last_br = &acc.Historic[len(acc.Historic)-1]
			last_br.Yesterday(prev_br.TotalByIDs, prev_br.TotalByCodes)
		}
		last_br.Add(asset_id, asset_code, amount)
	}

	// Reverse order
	l := len(acc.Historic) - 1
	for i := 0; i < l/2; i++ {
		a := acc.Historic[i]
		b := acc.Historic[l-i]
		acc.Historic[i] = b
		acc.Historic[l-i] = a
	}

	return nil
}

func (acc *Account) LoadBalanceAt(date LDate) error {
	rows, err := DB.Query("SELECT `AssetID`, `AssetCode`, SUM(`Amount`) FROM `movements_view` WHERE `AccountID` = ? AND `MovementDate` <= ? AND `MovementStatus` != 'C' GROUP BY `AccountID`, `AssetID`", acc.ID, date)

	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		Log.WarningF("Error when loading balance: %#v", err)
		return err
	}
	defer rows.Close()
	acc.LocalBalanceIDs = make(map[uuid.UUID]int64)
	acc.LocalBalanceCodes = make(map[string]int64)
	for rows.Next() {
		var asset_id uuid.UUID
		var asset_code string
		var amount int64
		err = rows.Scan(&asset_id, &asset_code, &amount)
		if err != nil {
			Log.WarningF("Error when loading balance for account %s: %#v", acc.ID.String(), err)
			return err
		}
		// Do not include zeros
		if amount == 0 {
			continue
		}

		acc.LocalBalanceIDs[asset_id] = amount
		acc.LocalBalanceCodes[asset_code] = amount
	}

	return nil
}

func AccountTree(input []Account) Account {
	// Add an empty/dummy account with null id to make things easier
	root := Account{}
	mask := make(map[uuid.UUID]bool)
	// Make the tree
	account_tree_body(&root, input, mask)
	account_tree_sum(&root)
	return root
}

func AccountList(root Account) []Account {
	ans := make([]Account, 1)
	ans[0] = root
	for _, child := range root.Children {
		buf := AccountList(child)
		ans = append(ans, buf...)
	}
	ans[0].Children = nil
	return ans
}

func account_tree_body(root *Account, input []Account, mask map[uuid.UUID]bool) {
	root.Children = make([]Account, 0)
	mask[root.ID] = true
	for _, account := range input {
		if uuid.Equal(root.ID, account.ParentID) && !uuid.Equal(root.ID, account.ID) {
			if !mask[account.ID] {
				account_tree_body(&account, input, mask)
			}
			root.Children = append(root.Children, account)
		}
	}
}

func account_tree_sum(root *Account) {
	// Copy stuff
	root.TotalBalanceIDs = make(map[uuid.UUID]int64)
	root.TotalBalanceCodes = make(map[string]int64)
	for id, val := range root.LocalBalanceIDs {
		root.TotalBalanceIDs[id] += val
	}
	for id, val := range root.LocalBalanceCodes {
		root.TotalBalanceCodes[id] += val
	}

	// Sum stuff
	for i := range root.Children {
		acc := &root.Children[i]
		account_tree_sum(acc)
		for id, val := range acc.TotalBalanceIDs {
			root.TotalBalanceIDs[id] += val
		}
		for code, val := range acc.TotalBalanceCodes {
			root.TotalBalanceCodes[code] += val
		}
	}
}

func (this AccountsDBConn) InBook(book_id uuid.UUID) ([]Account, error) {
	accounts := make([]Account, 0)
	rows, err := DB.Query("SELECT `ID`, `ParentID`, `Name`, `BookID` FROM `accounts` WHERE `BookID` = ? ORDER BY `Name`", book_id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading accounts: %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		account := Account{}
		err = rows.Scan(
			&account.ID,
			&account.ParentID,
			&account.Name,
			&account.BookID)
		if err != nil {
			Log.WarningF("Error when loading account %s: %#v", account.ID.String(), err)
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (this AccountsDBConn) ByID(acc_id uuid.UUID) (Account, error) {
	account := Account{}

	err := DB.QueryRow("SELECT `ID`, `ParentID`, `Name`, `BookID` FROM `accounts` WHERE `ID` = ? ORDER BY `Name`", acc_id).Scan(
		&account.ID,
		&account.ParentID,
		&account.Name,
		&account.BookID)
	if err == sql.ErrNoRows {
		return account, err
	}
	if err != nil {
		Log.WarningF("Error when loading account %s: %#v", acc_id.String(), err)
		return account, err
	}
	return account, nil
}

func (this AccountsDBConn) Set(account *Account) error {
	// If the ID is nil, assume it is a new account and give it a new ID
	if account.ID.IsNil() {
		account.ID = uuid.NewV4()
	}
	err := account.Validate()
	if err != nil {
		return err
	}
	_, err = DB.Exec("REPLACE INTO `accounts` VALUE (?, ?, ?, ?)",
		account.ID,
		account.ParentID,
		account.Name,
		account.BookID)
	if err != nil {
		Log.WarningF("Error when updating account %s %s: %#v", account.ID, account.Name, err)
		return FixError(err)
	}
	return nil
}
