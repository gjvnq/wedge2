package wedge

import (
	"time"

	uuid "github.com/gjvnq/go.uuid"
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
}
