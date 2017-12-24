package wedge

import (
	"github.com/gjvnq/go.uuid"
)

type Asset struct {
	ID     uuid.UUID `json:"id"`
	BookID uuid.UUID `json:"book_id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
}
