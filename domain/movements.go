package wedge

import (
	"database/sql"
	"strings"

	"github.com/gjvnq/go.uuid"
)

type MovementsDBConn struct{}

var Movements MovementsDBConn

type Movement struct {
	ID            uuid.UUID `json:"id"`
	BookID        uuid.UUID `json:"book_id"`
	AccountID     uuid.UUID `json:"account_id"`
	AssetID       uuid.UUID `json:"asset_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Amount        int64     `json:"amount"`
	Status        string    `json:"status"`
	LocalDate     LDate     `json:"local_date"`
	Notes         string    `json:"notes"`
	Tags          []string  `json:"tags"`
}

type MovementExtended struct {
	Movement
	TransactionName string `json:"transaction_name"`
	TransactionDate LDate  `json:"transaction_date"`
	AssetCode       string `json:"asset_code"`
}

func (mov Movement) GetID() uuid.UUID {
	return mov.ID
}

func (mov Movement) GetBookID() uuid.UUID {
	return mov.BookID
}

func (this MovementsDBConn) InAccountAndBook(acc_id, book_id uuid.UUID) ([]MovementExtended, error) {
	movements := make([]MovementExtended, 0)
	rows, err := DB.Query("SELECT `MovementID`, `BookID`, `AccountID`, `AssetID`, `TransactionID`, `Amount`, `MovementStatus`, `MovementDate`, `TransactionName`, `TransactionDate`, `AssetCode`, `Tags` FROM `movements_view` WHERE `AccountID` = ? AND `BookID` = ? ORDER BY `MovementDate` DESC, `TransactionName` ASC", acc_id, book_id)
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
		str := ""
		err = rows.Scan(
			&movement.ID,
			&movement.BookID,
			&movement.AccountID,
			&movement.AssetID,
			&movement.TransactionID,
			&movement.Amount,
			&movement.Status,
			&movement.LocalDate,
			&movement.TransactionName,
			&movement.TransactionDate,
			&movement.AssetCode,
			&str)
		movement.Tags = strings.Split(str, ",")
		if err != nil {
			Log.WarningF("Error when loading movement %s: %#v", movement.ID.String(), err)
			return nil, err
		}
		movements = append(movements, movement)
	}
	return movements, nil
}
