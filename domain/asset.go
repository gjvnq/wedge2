package wedge

import (
	"time"

	"github.com/satori/go.uuid"
)

type Asset struct {
	ID     uuid.UUID `json:"id"`
	BookID uuid.UUID `json:"book_id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Places int       `json:"places"`
	Fmt    string    `json:"fmt"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book   Book         `json:"book,omitempty"`
	Values []AssetValue `json:"values,omitempty"`
}
