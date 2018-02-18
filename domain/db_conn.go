package wedge

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gjvnq/go-logger"
)

var Log *logger.Logger
var DB *sql.DB

func FixError(err error) error {
	if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
		Log.Debug("[duplicate entry]", err)
		return errors.New("duplicate entry")
	}
	return err
}

func PrepareStatments() {
	var err error
	var sql string

	sql = "SELECT\n"
	sql += "`movements`.`ID` AS `MovementID`,\n"
	sql += "`movements`.`AccountID` AS `AccountID`,\n"
	sql += "`movements`.`AssetID` AS `AssetID`,\n"
	sql += "`movements`.`TransactionID` AS `TransactionID`,\n"
	sql += "`movements`.`Amount` AS `Amount`,\n"
	sql += "`movements`.`Status` AS `MovementStatus`,\n"
	sql += "`movements`.`LocalDate` AS `MovementDate`,\n"
	sql += "`movements`.`Notes` AS `MovementNotes`,\n"
	sql += "COALESCE(GROUP_CONCAT(`tags`.`Tag` SEPARATOR ','), '') AS `Tags`\n"
	sql += "FROM `movements`\n"
	sql += "LEFT JOIN `transactions` ON(`movements`.`TransactionID` = `transactions`.`ID`)\n"
	sql += "LEFT JOIN `assets` ON(`movements`.`AssetID` = `assets`.`ID`)\n"
	sql += "LEFT JOIN `tags` ON(`movements`.`ID` = `tags`.`ItemID`)\n"
	sql += "LEFT JOIN `accounts` ON(`movements`.`AccountID` = `accounts`.`ID`)\n"
	sql += "WHERE `movements`.`TransactionID` = ? "
	sql += "GROUP BY `movements`.`ID`"
	fillMovementsStmt, err = DB.Prepare(sql)
	if err != nil {
		Log.Fatal(err)
	}
}
