package wedge

import (
	"github.com/gjvnq/go.uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LocalDate LDate     `json:"local_date"`
	BookID    uuid.UUID `json:"book_id"`
	// Associations
	Movements []Movement `json:"movements,omitempty"`
	Items     []Item     `json:"items,omitempty"`
}
