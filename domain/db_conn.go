package wedge

import (
	"database/sql"

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
	rows, err := DB.Query("SELECT `ID`, `BookID`, `Name`, `Code`, `Places` FROM `assets` WHERE `BookID` = ?", book_id)
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
			&asset.Code,
			&asset.Places)
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
	err := DB.QueryRow("SELECT `ID`, `BookID`, `Name`, `Code`, `Places` FROM `assets` WHERE `ID` = ?", asset_id).Scan(
		&asset.ID,
		&asset.BookID,
		&asset.Name,
		&asset.Code,
		&asset.Places)
	if err == sql.ErrNoRows {
		return Asset{}, err, true
	}
	if err != nil {
		Log.WarningF("Error when loading asset %s: %#v", asset.ID.String(), err)
		return Asset{}, err, false
	}
	return asset, nil, false
}

func Assets_Update(asset *Asset) error {
	_, err := DB.Exec("UPDATE `assets` SET `BookID` = ?, `Name` = ?, `Code` = ?, `Places` = ? WHERE `ID` = ?",
		asset.BookID,
		asset.Name,
		asset.Code,
		asset.Places,
		asset.ID)
	if err != nil {
		Log.WarningF("Error when updating asset %s: %#v", asset.ID.String(), err)
		return err
	}
	return nil
}

func Assets_Insert(asset *Asset) error {
	if asset.ID.IsNil() {
		asset.ID = uuid.NewV4()
	}
	_, err := DB.Exec("INSERT INTO `assets` (`ID`, `BookID`, `Name`, `Code`, `Places`) VALUES (?, ?, ?, ?, ?);",
		asset.ID,
		asset.BookID,
		asset.Name,
		asset.Code,
		asset.Places)
	if err != nil {
		Log.WarningF("Error when inserting asset %+v: %#v", asset, err)
		return err
	}
	return nil
}
