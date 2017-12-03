package wedge

import (
	"fmt"
	"time"

	"github.com/gjvnq/go.uuid"
)

type Account struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
	BookID   uuid.UUID `json:"book_id"`
	ParentID uuid.UUID `json:"parent_id"`
	Name     string    `json:"name"`
	// Date Stuff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Associations
	Book     Book      `json:"book,omitempty"`
	Children []Account `json:"children"`
}

func (acc Account) String() string {
	return fmt.Sprintf("<%s %s %d>", acc.ID, acc.Name, len(acc.Children))
}

func AccountTree(input []Account) Account {
	// Add an empty/dummy account with null id to make things easier
	root := Account{}
	mask := make(map[uuid.UUID]bool)
	// Make the tree
	account_tree_body(&root, input, mask)
	return root
}

func account_tree_body(root *Account, input []Account, mask map[uuid.UUID]bool) {
	root.Children = make([]Account, 0)
	mask[root.ID] = true
	for _, account := range input {
		if uuid.Equal(root.ID, account.ParentID) && !uuid.Equal(root.ID, account.ID) {
			if !mask[account.ID] {
				account_tree_body(&account, input, mask)
			}
			root.Children = append(root.Children, account)
		}
	}
}
