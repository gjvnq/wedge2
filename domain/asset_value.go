package wedge

import (
	"time"

	"github.com/satori/go.uuid"
)

type AssetValue struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	AssetID   uuid.UUID `json:"asset_id"`
	BaseID    uuid.UUID `json:"base_id"`
	Value     float64   `json:"value"`
	LocalDate LDate     `json:"local_date"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book  Book  `json:"book,omitempty"`
	Asset Asset `json:"asset,omitempty"`
	Base  Asset `json:"base,omitempty"`
}
