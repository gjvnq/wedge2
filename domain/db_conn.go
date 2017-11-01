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
