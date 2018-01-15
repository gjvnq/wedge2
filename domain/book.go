package wedge

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"

	"github.com/gjvnq/go.uuid"
)

type BooksDBConn struct{}

var Books BooksDBConn

type Book struct {
	ID       uuid.UUID
	Name     string
	Password []byte
}

func (book *Book) SetPassword(password string) error {
	var err error
	book.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return err
}

func (book *Book) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(book.Password, []byte(password))
	return err == nil
}

func (book *Book) Redact() {
	book.Password = nil
}

func (this BooksDBConn) GetByID(id uuid.UUID) (Book, error, bool) {
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

func (this BooksDBConn) All(redact bool) ([]Book, error) {
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
