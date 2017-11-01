package wedge

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/gjvnq/go.uuid"
)

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
