package wedge

import "github.com/satori/go.uuid"

type Book struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
	Name string
	Password []byte
}