package wedge

import (
	"time"

	"github.com/gjvnq/go.uuid"
)

type Item struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key"`
	AssetID       uuid.UUID `json:"asset_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Name          string    `json:"name"`
	UnitCost      int64     `json:"unit_cost"`
	Qty           float64   `json:"quantity"`
	TotalCost     int64     `json:"total_cost"`
	PeriodEnd     LDate     `json:"period_end"`
	PeriodStart   LDate     `json:"period_start"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Asset       Asset       `json:"asset,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
}
