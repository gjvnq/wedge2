package wedge

import (
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
}
