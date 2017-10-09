package wedge

import (
	"github.com/satori/go.uuid"
	"time"
)

type Account struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
	BookID uuid.UUID `json:"book_id"`
	ParentID uuid.UUID `json:"parent_id"`
	Name string `json:"name"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book Book `json:"book,omitempty"`
}