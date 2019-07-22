package ausholiday

import (
	"fmt"
	"time"
)

// Days how many days shared between two date
func Days(from, to time.Time) int {
	ft := from.Truncate(time.Hour * 24)
	tt := to.Truncate(time.Hour * 24) // add 24 hours, to count last day inclusively

	return int(tt.Sub(ft).Hours() / 24.0)
}

// IsWeekend is it a weekend day
func IsWeekend(t time.Time) bool {
	day := t.Weekday()

	if day == 0 || day == 6 {
		return true
	}

	return false
}

// BusinessDays how many business days between two date in state
func BusinessDays(state State, from, to time.Time, holidays ...*Holiday) int {
	if holidays == nil {
		holidays = xholidays
	}

	// not using .Truncate because it doesnt truncate by timezone
	// force using canberra time (+10:00) to match the csv file time
	loc, _ := time.LoadLocation("Australia/Canberra")
	it, _ := time.ParseInLocation(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+10:00", from.Year(), int(from.Month()), from.Day()), loc)
	tt, _ := time.ParseInLocation(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+10:00", to.Year(), int(to.Month()), to.Day()), loc)

	// same day
	if it.Equal(tt) {
		return 0
	}
	// invalid time, to is before from
	if tt.Before(it) {
		return -1
	}

	count := 0
	for it.Before(tt) {
		it = it.AddDate(0, 0, 1)

		if IsHoliday(state, it, holidays...) || IsWeekend(it) {
			// fmt.Printf("   x > %+v  %d\n", it, count)
		} else {
			count++
			// fmt.Printf("     > %+v  %d\n", it, count)
		}
	}

	return count
}

// BusinessDaysAny how many business days between two date in any state
func BusinessDaysAny(from, to time.Time, holidays ...*Holiday) int {
	return BusinessDays("any", from, to, holidays...)
}

// OffDays how many holiday days between two date for state
func OffDays(state State, from, to time.Time, holidays ...*Holiday) int {
	if holidays == nil {
		holidays = xholidays
	}

	// not using .Truncate because it doesnt truncate by timezone
	// force using canberra time (+10:00) to match the csv file time
	loc, _ := time.LoadLocation("Australia/Canberra")
	it, _ := time.ParseInLocation(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+10:00", from.Year(), int(from.Month()), from.Day()), loc)
	tt, _ := time.ParseInLocation(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+10:00", to.Year(), int(to.Month()), to.Day()), loc)

	// same day
	if it.Equal(tt) {
		return 0
	}
	// invalid time, to is before from
	if tt.Before(it) {
		return -1
	}

	count := 0
	for it.Before(tt) {
		it = it.AddDate(0, 0, 1)

		if IsHoliday(state, it, holidays...) || IsWeekend(it) {
			count++
			// fmt.Printf("   x > %+v  %d\n", it, count)
		} else {
			// fmt.Printf("     > %+v  %d\n", it, count)
		}
	}

	return count
}

// OffDaysAny how many holiday days between two date for any state
func OffDaysAny(from, to time.Time, holidays ...*Holiday) int {
	return OffDays("any", from, to, holidays...)
}

// IsHoliday is it a holiday in State
func IsHoliday(state State, t time.Time, holidays ...*Holiday) bool {
	if holidays == nil {
		holidays = xholidays
	}

	iymd := fmt.Sprintf("%d%02d%02d", t.Year(), int(t.Month()), t.Day())

	for _, holiday := range holidays {
		// fmt.Printf("           %+v  ~  %+v   %+v \n", iymd, holiday.YMD, holiday.Date)

		// match any state
		if holiday.YMD == iymd && state == "any" {
			return true
		}

		if holiday.YMD == iymd && holiday.Jurisdiction == state {
			return true
		}
	}

	return false
}

// IsHolidayAny is it a holiday in any state
func IsHolidayAny(t time.Time, holidays ...*Holiday) bool {
	return IsHoliday("any", t, holidays...)
}
