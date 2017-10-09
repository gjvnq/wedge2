package wedge

import (
	"github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
	Name string	`json:"name"`
	LocalDate LDate	`json:"local_date"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Movements []Movement `json:"movements,omitempty"`
	Itens []Item `json:"itens,omitempty"`
}