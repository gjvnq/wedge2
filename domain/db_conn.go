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
