package wedge

import (
	"github.com/satori/go.uuid"
	"time"
)

type Asset struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
	BookID uuid.UUID `json:"book_id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Fmt string `json:"fmt"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book Book `json:"book,omitempty"`
	Values []AssetValue `json:"values,omitempty"`
}