package wedge

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gjvnq/go.uuid"
)

type TagsDBConn struct{}

var Tags TagsDBConn

type TagI interface {
	GetID() uuid.UUID
	GetBookID() uuid.UUID
}

func (this TagsDBConn) Set(obj TagI, tags []string) error {
	// Start transaction
	tx, err := DB.Begin()
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when creating transaction: %#v", err)
		return FixError(err)
	}

	err = this.SetTX(tx, obj, tags)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when finishing db transaction: %#v", err)
		return FixError(err)
	}
	return nil
}

func (this TagsDBConn) SetTX(tx *sql.Tx, obj TagI, tags []string) error {
	var err error

	// Check for nil
	if obj.GetID().IsNil() {
		return errors.New("id must not be nil")
	}
	if tags == nil {
		tags = make([]string, 0)
	}

	// First delete ALL tags for this object
	_, err = DB.Exec("DELETE FROM `tags` WHERE `ItemID` = ?", obj.GetID())

	// Avoid stupid work
	if len(tags) == 0 {
		err = tx.Commit()
		if err != nil {
			return FixError(err)
		}
		Log.WarningF("Error when finishing db transaction: %#v", err)
		return nil
	}

	// Now reinsert tags
	sql_str := "INSERT INTO `tags` (`ItemID`, `BookID`, `Tag`) VALUES "
	vals := make([]interface{}, 0)
	for i, tag := range tags {
		if i == 0 {
			sql_str += "(?, ?, ?)"
		} else {
			sql_str += ", (?, ?, ?)"
		}
		vals = append(vals, obj.GetID(), obj.GetBookID(), strings.TrimSpace(tag))
	}
	stmt, _ := DB.Prepare(sql_str)
	_, err = stmt.Exec(vals...)
	if err != nil {
		tx.Rollback()
		Log.WarningF("Error when saving new tags: %#v", err)
		return FixError(err)
	}

	return nil
}

func (this TagsDBConn) Get(obj TagI) ([]string, error) {
	tags := make([]string, 0)
	rows, err := DB.Query("SELECT `Tag` FROM `tags` WHERE `ItemID` = ?", obj.GetID())
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		Log.WarningF("Error when loading tags: %#v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			Log.WarningF("Error when loading tags for %s: %#v", obj.GetID().String(), err)
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
