package wedge

import (
	"database/sql"
	"errors"

	"github.com/gjvnq/go.uuid"
)

type AssetsDBConn struct{}

var Assets AssetsDBConn

type Asset struct {
	ID     uuid.UUID `json:"id"`
	BookID uuid.UUID `json:"book_id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
}

func (this AssetsDBConn) InBook(book_id uuid.UUID) ([]Asset, error) {
	assets := make([]Asset, 0)
	rows, err := DB.Query("SELECT `ID`, `BookID`, `Name`, `Code` FROM `assets` WHERE `BookID` = ? ORDER BY CHAR_LENGTH(`Code`), `Code`", book_id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading assets: %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		asset := Asset{}
		err = rows.Scan(
			&asset.ID,
			&asset.BookID,
			&asset.Name,
			&asset.Code)
		if err != nil {
			Log.WarningF("Error when loading asset %s: %#v", asset.ID.String(), err)
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (this AssetsDBConn) GetById(asset_id uuid.UUID) (Asset, error, bool) {
	asset := Asset{}
	err := DB.QueryRow("SELECT `ID`, `BookID`, `Name`, `Code` FROM `assets` WHERE `ID` = ?", asset_id).Scan(
		&asset.ID,
		&asset.BookID,
		&asset.Name,
		&asset.Code)
	if err == sql.ErrNoRows {
		return Asset{}, err, true
	}
	if err != nil {
		Log.WarningF("Error when loading asset %s: %#v", asset.ID.String(), err)
		return Asset{}, err, false
	}
	return asset, nil, false
}

func (this AssetsDBConn) Set(asset *Asset) error {
	tx, err := DB.Begin()

	// If the ID is nil, assume it is a new asset and give it a new ID
	if asset.ID.IsNil() {
		asset.ID = uuid.NewV4()
		// Ensure there is no asset with the same code
		counter := -1
		err = tx.QueryRow("SELECT COUNT(*) FROM `assets` WHERE `Code` = ? AND `BookID` = ?",
			asset.Code,
			asset.BookID).Scan(&counter)
		if err != nil {
			Log.WarningF("Error when counting assets with code %s in book %s: %#v", asset.Code, asset.BookID, err)
			return FixError(err)
		}
		if counter != 0 {
			tx.Rollback()
			return errors.New("duplicate entry")
		}
	}
	_, err = tx.Exec("REPLACE INTO `assets` (`Id`, `Name`, `Code`, `BookID`) VALUE (?, ?, ?, ?)",
		asset.ID,
		asset.Name,
		asset.Code,
		asset.BookID)
	if err != nil {
		Log.WarningF("Error when updating asset %s %s: %#v", asset.ID, asset.Name, err)
		return FixError(err)
	}
	err = tx.Commit()
	if err != nil {
		Log.WarningF("Error when updating asset %s %s: %#v", asset.ID, asset.Name, err)
		return FixError(err)
	}
	return nil
}
