package wedge

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gjvnq/go.uuid"
)

type AccountsDBConn struct{}

var Accounts AccountsDBConn

type Account struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
	BookID   uuid.UUID `json:"book_id"`
	ParentID uuid.UUID `json:"parent_id"`
	Name     string    `json:"name"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book     Book      `json:"book,omitempty"`
	Children []Account `json:"children"`
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

func AccountTree(input []Account) Account {
	// Add an empty/dummy account with null id to make things easier
	root := Account{}
	mask := make(map[uuid.UUID]bool)
	// Make the tree
	account_tree_body(&root, input, mask)
	return root
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
