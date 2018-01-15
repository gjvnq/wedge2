package wedge

import (
	"database/sql"

	"github.com/gjvnq/go.uuid"
)

type TransactionsDBConn struct{}

var Transactions TransactionsDBConn

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LocalDate LDate     `json:"local_date"`
	BookID    uuid.UUID `json:"book_id"`
	// Associations
	Movements []Movement          `json:"movements"`
	Items     []Item              `json:"items"`
	Totals    map[uuid.UUID]int64 `json:"totals"`
}

func (tr *Transaction) Init() {
	if tr.Totals == nil {
		tr.Totals = make(map[uuid.UUID]int64)
	}
}

func (tr *Transaction) ComputeTotals() {
	tr.Totals = make(map[uuid.UUID]int64)
	for _, mov := range tr.Movements {
		tr.Totals[mov.AssetID] += mov.Amount
	}
}

func (this TransactionsDBConn) InBook(book_id uuid.UUID) ([]Transaction, error) {
	transactions := make([]Transaction, 0)
	rows, err := DB.Query("SELECT `ID`, `Name`, `LocalDate`, `BookID` FROM `transactions` WHERE `BookID` = ? ORDER BY `LocalDate`, `Name`", book_id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading transactions: %#v", err)
		return nil, err
	}
	defer rows.Close()
	// Load basic stuff
	for rows.Next() {
		transaction := Transaction{}
		err = rows.Scan(
			&transaction.ID,
			&transaction.Name,
			&transaction.LocalDate,
			&transaction.BookID)
		if err != nil {
			Log.WarningF("Error when loading transaction %s: %#v", transaction.ID.String(), err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	// Fill in stuff
	for i := 0; i < len(transactions); i++ {
		Transactions.FillMovements(&transactions[i])
		Transactions.FillItems(&transactions[i])
	}
	return transactions, nil
}

func (this TransactionsDBConn) GetByID(tr_id uuid.UUID) (Transaction, error, bool) {
	return Transactions.GetByIDAdv(tr_id, true, true)
}

func (this TransactionsDBConn) GetHeadByID(tr_id uuid.UUID) (Transaction, error, bool) {
	return Transactions.GetByIDAdv(tr_id, false, false)
}

func (this TransactionsDBConn) GetByIDAdv(tr_id uuid.UUID, include_movements bool, include_itens bool) (Transaction, error, bool) {
	tr := Transaction{}
	err := DB.QueryRow("SELECT `ID`, `Name`, `LocalDate`, `BookID` FROM `transactions` WHERE `ID` = ? ORDER BY `LocalDate`, `Name`", tr_id).Scan(
		&tr.ID,
		&tr.Name,
		&tr.LocalDate,
		&tr.BookID)
	if err == sql.ErrNoRows {
		return tr, nil, true
	}
	if err != nil {
		Log.WarningF("Error when loading transactions: %#v", err)
		return tr, err, false
	}
	if include_movements {
		if err = Transactions.FillMovements(&tr); err != nil {
			return tr, err, false
		}
	}
	if include_itens {
		if err = Transactions.FillItems(&tr); err != nil {
			return tr, err, false
		}
	}
	return tr, nil, false
}

func (this TransactionsDBConn) FillMovements(transaction *Transaction) error {
	transaction.Movements = make([]Movement, 0)
	rows, err := DB.Query("SELECT `ID`, `AccountID`, `AssetID`, `TransactionID`, `Amount`, `Status`, `LocalDate`, `Notes` FROM `movements` WHERE `TransactionID` = ? ORDER BY `LocalDate`, `Amount`", transaction.ID)
	if err == sql.ErrNoRows {
		return FixError(err)
	}
	if err != nil {
		Log.WarningF("Error when loading movements: %#v", err)
		return FixError(err)
	}
	defer rows.Close()
	for rows.Next() {
		movement := Movement{}
		err = rows.Scan(
			&movement.ID,
			&movement.AccountID,
			&movement.AssetID,
			&movement.TransactionID,
			&movement.Amount,
			&movement.Status,
			&movement.LocalDate,
			&movement.Notes)
		if err != nil {
			Log.WarningF("Error when loading movement %s: %#v", movement.ID.String(), err)
			return FixError(err)
		}
		transaction.Movements = append(transaction.Movements, movement)
	}
	transaction.ComputeTotals()
	return nil
}

func (this TransactionsDBConn) FillItems(transaction *Transaction) error {
	transaction.Items = make([]Item, 0)
	rows, err := DB.Query("SELECT `ID`, `AssetID`, `TransactionID`, `Name`, `UnitCost`, `Qty`, `TotalCost`, `PeriodEnd`, `PeriodStart` FROM `items` WHERE `TransactionID` = ? ORDER BY `Name`, `TotalCost`", transaction.ID)
	if err == sql.ErrNoRows {
		return FixError(err)
	}
	if err != nil {
		Log.WarningF("Error when loading items: %#v", err)
		return FixError(err)
	}
	defer rows.Close()
	for rows.Next() {
		item := Item{}
		err = rows.Scan(
			&item.ID,
			&item.AssetID,
			&item.TransactionID,
			&item.Name,
			&item.UnitCost,
			&item.Qty,
			&item.TotalCost,
			&item.PeriodStart,
			&item.PeriodEnd)
		if err != nil {
			Log.WarningF("Error when loading item %s: %#v", item.ID.String(), err)
			return FixError(err)
		}
		transaction.Items = append(transaction.Items, item)
	}
	return nil
}

func (this TransactionsDBConn) Set(transaction *Transaction) error {
	var err error
	// If the ID is nil, assume it is a new transaction and give it a new ID
	if transaction.ID.IsNil() {
		transaction.ID = uuid.NewV4()
	}
	tx, err := DB.Begin()
	if err != nil {
		Log.WarningF("Error when creating transaction: %#v", err)
		return FixError(err)
	}
	_, err = DB.Exec("REPLACE INTO `transactions` (`ID`, `Name`, `LocalDate`, `BookID`) VALUE (?, ?, ?, ?)",
		transaction.ID,
		transaction.Name,
		transaction.LocalDate,
		transaction.BookID)
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when creating transaction: %#v", err)
		return FixError(err)
	}
	_, err = tx.Exec("DELETE FROM `movements` WHERE `TransactionID` = ?",
		transaction.ID)
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when creating transaction: %#v", err)
		return FixError(err)
	}
	_, err = tx.Exec("DELETE FROM `items` WHERE `TransactionID` = ?",
		transaction.ID)
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when creating transaction: %#v", err)
		return FixError(err)
	}
	for i := 0; i < len(transaction.Movements); i++ {
		mov := &transaction.Movements[i]
		if mov.ID.IsNil() {
			mov.ID = uuid.NewV4()
		}
		mov.TransactionID = transaction.ID
		_, err = tx.Exec("INSERT INTO `movements` (`ID`, `AccountID`, `AssetID`, `TransactionID`, `Amount`, `Status`, `LocalDate`, `Notes`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			mov.ID,
			mov.AccountID,
			mov.AssetID,
			mov.TransactionID,
			mov.Amount,
			mov.Status,
			mov.LocalDate,
			mov.Notes)
		if err != nil {
			tx.Rollback()
			Log.WarningF("Error when creating transaction: %#v", err)
			return FixError(err)
		}
	}
	for i := 0; i < len(transaction.Items); i++ {
		item := &transaction.Items[i]
		if item.ID.IsNil() {
			item.ID = uuid.NewV4()
		}
		item.TransactionID = transaction.ID
		_, err = tx.Exec("INSERT INTO `items` (`ID`, `AssetID`, `TransactionID`, `Name`, `UnitCost`, `Qty`, `TotalCost`, `PeriodStart`, `PeriodEnd`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			item.ID,
			item.AssetID,
			item.TransactionID,
			item.Name,
			item.UnitCost,
			item.Qty,
			item.TotalCost,
			item.PeriodStart,
			item.PeriodEnd)
		if err != nil {
			tx.Rollback()
			Log.WarningF("Error when creating transaction: %#v", err)
			return FixError(err)
		}
	}
	Log.Debug("Commiting")
	err = tx.Commit()
	if err != nil {
		Log.WarningF("Error when updating transaction %s %s: %#v", transaction.ID, transaction.Name, err)
		return FixError(err)
	}

	return nil
}

func (this TransactionsDBConn) RmByID(tr_id uuid.UUID) error {
	var err error
	tx, err := DB.Begin()
	_, err = tx.Exec("DELETE FROM `transactions` WHERE `ID` = ?", tr_id)
	if err != nil {
		Log.WarningF("Error when deleting transaction head: %#v", err)
		return err
	}
	_, err = tx.Exec("DELETE FROM `movements` WHERE `TransactionID` = ?", tr_id)
	if err != nil {
		Log.WarningF("Error when deleting transaction movements: %#v", err)
		return err
	}
	_, err = tx.Exec("DELETE FROM `items` WHERE `TransactionID` = ?", tr_id)
	if err != nil {
		Log.WarningF("Error when deleting transaction items: %#v", err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		Log.WarningF("Error when deleting transaction: %#v", err)
		return err
	}
	return nil
}
