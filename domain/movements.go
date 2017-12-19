package wedge

import (
	"time"

	"github.com/gjvnq/go.uuid"
)

type Movement struct {
	ID            uuid.UUID `json:"id"`
	AccountID     uuid.UUID `json:"account_id"`
	AssetID       uuid.UUID `json:"asset_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Amount        int64     `json:"amount"`
	Status        string    `json:"status"`
	LocalDate     LDate     `json:"local_date"`
	Notes         string    `json:"notes"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Asset       Asset       `json:"asset,omitempty"`
	Account     Account     `json:"account,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
}
