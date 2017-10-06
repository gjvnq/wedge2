package main

import "github.com/satori/go.uuid"

type DBConn interface {
	// Misc
	Any(uuid string) (interface{}, error)
	ByTag(tag string) (map[uuid.UUID]string, error)
	IdsOfChildrenOfAccount(uuid string) ([]uuid.UUID, error)
	NewId() uuid.UUID
	SetTags(id uuid.UUID, tags []string) error
	// Movementation(account_id uuid.UUID, asset_id uuid.UUID, include_planned bool, start LocalDate, end LocalDate) (int, error)
	// Movementations(account_id uuid.UUID, include_planned bool, start LocalDate, end LocalDate) (map[uuid.UUID]int, error)
	// MovementationsHistory(account_id uuid.UUID, include_planned bool, start LocalDate, end LocalDate) (map[LocalDate]map[uuid.UUID]int, error)
	Balance(account_id uuid.UUID, asset_id uuid.UUID, include_planned bool, date LocalDate) (int, error)
	Balances(account_id uuid.UUID, include_planned bool, date LocalDate) (map[uuid.UUID]int, error)
	BalancesHistory(account_id uuid.UUID, include_planned bool, date LocalDate) (map[LocalDate]map[uuid.UUID]int, error)

	// Load
	Account(uuid string, recursive bool) (Account, error)
	Asset(uuid string) (Asset, error)
	Item(uuid string, recursive bool) (Item, error)
	Movement(uuid string, recursive bool) (Movement, error)
	Transaction(uuid string, recursive bool) (Transaction, error)

	// Save
	SetAccount(data Account) error
	SetAsset(data Asset) error
	SetItem(data Item) error
	SetMovement(data Movement) error
	SetTransaction(data Transaction) error

	// Del
	DelAccount(id uuid.UUID) error
	DelAsset(id uuid.UUID) error
	DelItem(id uuid.UUID) error
	DelMovement(id uuid.UUID) error
	DelTransaction(id uuid.UUID) error
}

// Should implement DBConn
type SQLImpl struct {}
// Implement Balance and Balances as wrappers arround BalancesHistory. The program may be slower to run but will be faster to develop