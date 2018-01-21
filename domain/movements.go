package wedge

import (
	"database/sql"

	"github.com/gjvnq/go.uuid"
)

type MovementsDBConn struct{}

var Movements MovementsDBConn

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

type MovementExtended struct {
	Movement
	TransactionName string `json:"transaction_name"`
	TransactionDate LDate  `json:"transaction_date"`
	AccountName     string `json:"account_name"`
	AssetName       string `json:"asset_name"`
}

func (this MovementsDBConn) InAccount(acc_id uuid.UUID) ([]MovementExtended, error) {
	movements := make([]MovementExtended, 0)
	rows, err := DB.Query("SELECT `MovementID`, `AccountID`, `AssetID`, `TransactionID`, `Amount`, `MovementStatus`, `MovementDate`, `TransactionName`, `TransactionDate`, `AccountName`, `AssetName` FROM `movements_view` WHERE `AccountID` = ? ORDER BY `MovementDate` DESC, `TransactionName` ASC", acc_id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading movements: %#v", err)
		return nil, err
	}
	defer rows.Close()
	// Load basic stuff
	for rows.Next() {
		movement := MovementExtended{}
		err = rows.Scan(
			&movement.ID,
			&movement.AccountID,
			&movement.AssetID,
			&movement.TransactionID,
			&movement.Amount,
			&movement.Status,
			&movement.LocalDate,
			&movement.TransactionName,
			&movement.AccountName,
			&movement.AssetName)
		if err != nil {
			Log.WarningF("Error when loading movement %s: %#v", movement.ID.String(), err)
			return nil, err
		}
		movements = append(movements, movement)
	}
	return movements, nil
}
