package wedge

import (
	"time"

	"github.com/satori/go.uuid"
)

type Item struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key"`
	AssetID       uuid.UUID `json:"asset_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Name          string    `json:"name"`
	GenericName   string    `json:"generic_name"`
	Unit          string    `json:"unit"`
	UnitCost      int       `json:"unit_cost"`
	Qty           float64   `json:"qty"`
	TotalCost     int       `json:"total_cost"`
	PeriodEnd     LDate     `json:"period_end"`
	PeriodStart   LDate     `json:"period_start"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Asset       Asset       `json:"asset,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
}
