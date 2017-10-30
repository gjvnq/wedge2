package wedge

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Book struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
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

func (book *Book) BeforeCreate(scope *gorm.Scope) error {
	fmt.Printf("%+v\n", scope)
	fmt.Println(scope.SetColumn("ID", uuid.NewV4()))
	fmt.Printf("%+v\n", scope)
	return nil
}
