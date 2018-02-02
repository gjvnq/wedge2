package wedge

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DAY_ANY = 0
const YEAR_ANY = 0
const MONTH_ANY = 0
const MONTH_JAN = 1
const MONTH_FEB = 2
const MONTH_MAR = 3
const MONTH_APR = 4
const MONTH_MAY = 5
const MONTH_JUN = 6
const MONTH_JUL = 7
const MONTH_AUG = 8
const MONTH_SEP = 9
const MONTH_OCT = 10
const MONTH_NOV = 11
const MONTH_DEC = 12

type LDate struct {
	year  int
	month int
	day   int
}

func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}
	if year%400 != 0 {
		return false
	}
	return true
}

func DaysInMonth(month int, year int) int {
	return DaysInMonthLeap(month, IsLeapYear(year))
}

func DaysInMonthLeap(month int, leap bool) int {
	switch month {
	case MONTH_ANY:
		return 31
	case MONTH_JAN:
		return 31
	case MONTH_FEB:
		if leap {
			return 29
		} else {
			return 28
		}
	case MONTH_MAR:
		return 31
	case MONTH_APR:
		return 30
	case MONTH_MAY:
		return 31
	case MONTH_JUN:
		return 30
	case MONTH_JUL:
		return 31
	case MONTH_AUG:
		return 31
	case MONTH_SEP:
		return 30
	case MONTH_OCT:
		return 31
	case MONTH_NOV:
		return 30
	case MONTH_DEC:
		return 31
	default:
		return -1
	}
}

func DaysInMonthYear(month int, year int) int {
	return DaysInMonthLeap(month, IsLeapYear(year))
}

func (ld LDate) month_day_to_year_day() int {
	if !ld.is_valid() {
		return -1
	}
	count := 0
	for i := 1; i <= ld.month; i++ {
		count += DaysInMonth(i, ld.year)
	}
	count += ld.day
	return count
}

func (ld LDate) is_valid() bool {
	max_days := DaysInMonth(ld.month, ld.year)
	return 0 <= ld.day && ld.day <= max_days
}

func (ld LDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", ld.year, ld.month, ld.day)
}

func (ld1 LDate) Equals(ld2 LDate) bool {
	return ld1.String() == ld2.String()
}

func (ld *LDate) Set(y, m, d int) {
	ld.year = y
	ld.month = m
	ld.day = d
}

func LDateNow() LDate {
	ld := LDate{}
	t := time.Now()
	ld.Set(t.Year(), int(t.Month()), t.Day())
	return ld
}

func Limit99(v int) int {
	if v < 0 {
		return 0
	}
	if v > 99 {
		return 99
	}
	return v
}

func (ld LDate) Value() (driver.Value, error) {
	return int64(int(1E4)*ld.year + int(1E2)*Limit99(ld.month) + Limit99(ld.day)), nil
}

func (ld *LDate) Scan(value interface{}) error {
	if value == nil {
		ld.year = 0
		ld.month = 0
		ld.day = 0
		return nil
	}
	if int32v, err := driver.Int32.ConvertValue(value); err == nil {
		v := 0
		switch iv := int32v.(type) {
		case int:
			v = int(iv)
		case int32:
			v = int(iv)
		case int64:
			v = int(iv)
		default:
			return errors.New("failed to scan LDate (#1)")
		}
		ld.year = v / int(1E4)
		v = v % int(1E4)
		ld.month = v / 1E2
		ld.day = v % int(1E2)
		if ld.year < 1582 && ld.year != 0 {
			return errors.New("year must be 1582 or after")
		}
		if ld.year > 9999 {
			return errors.New("year must be 9999 or earlier")
		}
		return nil
	}
	return errors.New("failed to scan LDate (#2)")
}

func (ld *LDate) UnmarshalJSON(raw_data []byte) error {
	var err error

	data := string(raw_data)
	data = strings.Replace(data, "\"", "", -1)
	numbers := strings.Split(data, "-")
	if len(numbers) != 3 {
		return errors.New("invalid format, it MUST be yyyy-mm-dd")
	}
	ld.year, err = strconv.Atoi(numbers[0])
	if err != nil {
		Log.Error(err)
		return errors.New("'" + numbers[0] + "' is not a valid integer")
	}
	if ld.year < 1582 && ld.year != 0 {
		return errors.New("year must be 1582 or after")
	}
	if ld.year > 9999 {
		return errors.New("year must be 9999 or earlier")
	}
	ld.month, err = strconv.Atoi(numbers[1])
	if err != nil {
		return errors.New("'" + numbers[1] + "' is not a valid integer")
	}
	if !(0 <= ld.month && ld.month <= 12) {
		return errors.New("month must be between 0 and 12")
	}
	ld.day, err = strconv.Atoi(numbers[2])
	if err != nil {
		return errors.New("'" + numbers[2] + "' is not a valid integer")
	}
	max_days := DaysInMonthYear(ld.month, ld.year)
	if !(0 <= ld.day && ld.day <= max_days) {
		return errors.New("day must be between 0 and " + strconv.Itoa(max_days))
	}

	return nil
}

func (ld LDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + ld.String() + "\""), nil
}
