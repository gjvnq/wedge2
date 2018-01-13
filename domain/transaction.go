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
	Movements []Movement          `json:"movements"`
	Items     []Item              `json:"items"`
	Totals    map[uuid.UUID]int64 `json:"totals"`
}

func (tr *Transaction) Init() {
	if tr.Totals == nil {
		tr.Totals = make(map[uuid.UUID]int64)
	}
}

func (tr *Transaction) ComputeTotals() {
	tr.Totals = make(map[uuid.UUID]int64)
	for _, mov := range tr.Movements {
		tr.Totals[mov.AssetID] += mov.Amount
	}
}
