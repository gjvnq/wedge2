package wedge

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gjvnq/go-logger"
	"github.com/gjvnq/go.uuid"
)

var Log *logger.Logger
var DB *sql.DB

func Books_GetByID(id uuid.UUID) (Book, error, bool) {
	book := Book{}
	err := DB.QueryRow("SELECT `ID`, `Name`, `Password` FROM `books` WHERE `ID` = ?", id).Scan(&book.ID, &book.Name, &book.Password)
	if err == sql.ErrNoRows {
		return Book{}, err, true
	}
	if err != nil {
		Log.WarningF("Error when loading book %s: %#v", book.ID.String(), err)
		return Book{}, err, false
	}
	return book, nil, false
}

func Books_All(redact bool) ([]Book, error) {
	books := make([]Book, 0)
	rows, err := DB.Query("SELECT `ID`, `Name`, `Password` FROM `books`")
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading books: %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.Password)
		if err != nil {
			Log.WarningF("Error when loading book %s: %#v", book.ID.String(), err)
			return nil, err
		}
		if redact {
			book.Redact()
		}
		books = append(books, book)
	}
	return books, nil
}

func Assets_InBook(book_id uuid.UUID) ([]Asset, error) {
	assets := make([]Asset, 0)
	rows, err := DB.Query("SELECT `ID`, `BookID`, `Name`, `Code` FROM `assets` WHERE `BookID` = ? ORDER BY CHAR_LENGTH(`Code`), `Code`", book_id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading assets: %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		asset := Asset{}
		err = rows.Scan(
			&asset.ID,
			&asset.BookID,
			&asset.Name,
			&asset.Code)
		if err != nil {
			Log.WarningF("Error when loading asset %s: %#v", asset.ID.String(), err)
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func Assets_GetById(asset_id uuid.UUID) (Asset, error, bool) {
	asset := Asset{}
	err := DB.QueryRow("SELECT `ID`, `BookID`, `Name`, `Code` FROM `assets` WHERE `ID` = ?", asset_id).Scan(
		&asset.ID,
		&asset.BookID,
		&asset.Name,
		&asset.Code)
	if err == sql.ErrNoRows {
		return Asset{}, err, true
	}
	if err != nil {
		Log.WarningF("Error when loading asset %s: %#v", asset.ID.String(), err)
		return Asset{}, err, false
	}
	return asset, nil, false
}

func Assets_Set(asset *Asset) error {
	tx, err := DB.Begin()

	// If the ID is nil, assume it is a new asset and give it a new ID
	if asset.ID.IsNil() {
		asset.ID = uuid.NewV4()
		// Ensure there is no asset with the same code
		counter := -1
		err = tx.QueryRow("SELECT COUNT(*) FROM `assets` WHERE `Code` = ? AND `BookID` = ?",
			asset.Code,
			asset.BookID).Scan(&counter)
		if err != nil {
			Log.WarningF("Error when counting assets with code %s in book %s: %#v", asset.Code, asset.BookID, err)
			return FixError(err)
		}
		if counter != 0 {
			tx.Rollback()
			return errors.New("duplicate entry")
		}
	}
	_, err = tx.Exec("REPLACE INTO `assets` (`Id`, `Name`, `Code`, `BookID`) VALUE (?, ?, ?, ?)",
		asset.ID,
		asset.Name,
		asset.Code,
		asset.BookID)
	if err != nil {
		Log.WarningF("Error when updating asset %s %s: %#v", asset.ID, asset.Name, err)
		return FixError(err)
	}
	err = tx.Commit()
	if err != nil {
		Log.WarningF("Error when updating asset %s %s: %#v", asset.ID, asset.Name, err)
		return FixError(err)
	}
	return nil
}

func Accounts_InBook(book_id uuid.UUID) ([]Account, error) {
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

func Accounts_Set(account *Account) error {
	// If the ID is nil, assume it is a new account and give it a new ID
	if account.ID.IsNil() {
		account.ID = uuid.NewV4()
	}
	_, err := DB.Exec("REPLACE INTO `accounts` VALUE (?, ?, ?, ?)",
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

func Transactions_InBook(book_id uuid.UUID) ([]Transaction, error) {
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
		Transactions_FillMovements(&transactions[i])
		Transactions_FillItems(&transactions[i])
	}
	return transactions, nil
}

func Transactions_GetByID(tr_id uuid.UUID) (Transaction, error, bool) {
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
	if err = Transactions_FillMovements(&tr); err != nil {
		return tr, err, false
	}
	if err = Transactions_FillItems(&tr); err != nil {
		return tr, err, false
	}
	return tr, nil, false
}

func Transactions_FillMovements(transaction *Transaction) error {
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

func Transactions_FillItems(transaction *Transaction) error {
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

func Transactions_Set(transaction *Transaction) error {
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

func FixError(err error) error {
	if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
		Log.Debug("[duplicate entry]", err)
		return errors.New("duplicate entry")
	}
	return err
}
