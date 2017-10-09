package wedge

import "fmt"

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
	year int
	month int
	month_day int
	year_day int
}

func IsLeapYear(year int) bool {
	if year % 4 != 0 {
		return false
	}
	if year % 100 != 0 {
		return true
	}
	if year % 400 != 0 {
		return false
	}
	return true
}

func DaysInMonth(month int, year int) int {
	return DaysInMonthLeap(month, IsLeapYear(year))
}

func DaysInMonthLeap(month int, leap bool) int {
	switch (month) {
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

func (ld LDate) month_day_to_year_day() int {
	if !ld.is_valid() {
		return -1
	}
	count := 0
	for i := 1; i <= ld.month; i++ {
		count += DaysInMonth(i, ld.year)
	}
	count += ld.month_day
	return count
}

func (ld LDate) is_valid() bool {
	max_days := DaysInMonth(ld.month, ld.year)
	return 0 <= ld.month_day && ld.month_day <= max_days
}

func (ld LDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02dLD", ld.year, ld.month, ld.month_day)
}