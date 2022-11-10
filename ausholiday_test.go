package ausholiday

import (
	"os"
	"strings"
	"testing"
	"time"
)

var loc *time.Location
var d1 time.Time
var d2 time.Time
var ausHDay *AusHoliday

func TestDays(t *testing.T) {
	var t1, t2 time.Time

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 17)
	t.Run("2019-04-17,17", testDaysFunc(ausHDay, t1, t2, 0))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 10)
	t.Run("2019-04-17,17", testDaysFunc(ausHDay, t1, t2, -7))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 18)
	t.Run("2019-04-17,18", testDaysFunc(ausHDay, t1, t2, 1))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 22)
	t.Run("2019-04-17,22", testDaysFunc(ausHDay, t1, t2, 5))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 24)
	t.Run("2019-04-17,24", testDaysFunc(ausHDay, t1, t2, 7))
}
func testDaysFunc(hday *AusHoliday, t1, t2 time.Time, expect int) func(t *testing.T) {
	return func(t *testing.T) {
		days := hday.Days(t1, t2)
		if days != expect {
			t.Errorf("Days was incorrect, got: %d, want: %d", days, expect)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	var t1 time.Time

	t1 = ezTime(2019, 4, 18)
	t.Run("2019-04-18", testIsWeekendFunc(ausHDay, t1, false))

	t1 = ezTime(2019, 4, 19)
	t.Run("2019-04-19", testIsWeekendFunc(ausHDay, t1, false))

	t1 = ezTime(2019, 4, 20)
	t.Run("2019-04-20", testIsWeekendFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 21)
	t.Run("2019-04-21", testIsWeekendFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 22)
	t.Run("2019-04-22", testIsWeekendFunc(ausHDay, t1, false))

	t1 = ezTime(2019, 4, 23)
	t.Run("2019-04-23", testIsWeekendFunc(ausHDay, t1, false))
}
func testIsWeekendFunc(hday *AusHoliday, t1 time.Time, expect bool) func(t *testing.T) {
	return func(t *testing.T) {
		b := hday.IsWeekend(t1)
		if b != expect {
			t.Errorf("IsWeekend was incorrect, got: %t, want: %t", b, expect)
		}
	}
}

func TestIsHoliday(t *testing.T) {
	var t1 time.Time

	t1 = ezTime(2019, 4, 18)
	t.Run("2019-04-18", testIsHolidayFunc(ausHDay, t1, false))

	t1 = ezTime(2019, 4, 19)
	t.Run("2019-04-19", testIsHolidayFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 20)
	t.Run("2019-04-20", testIsHolidayFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 21)
	t.Run("2019-04-21", testIsHolidayFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 22)
	t.Run("2019-04-22", testIsHolidayFunc(ausHDay, t1, true))

	t1 = ezTime(2019, 4, 23)
	t.Run("2019-04-23", testIsHolidayFunc(ausHDay, t1, false))
}
func testIsHolidayFunc(hday *AusHoliday, t1 time.Time, expect bool) func(t *testing.T) {
	return func(t *testing.T) {
		b := hday.IsHoliday(ACT, t1)
		if b != expect {
			t.Errorf("IsHoliday was incorrect, got: %t, want: %t", b, expect)
		}
	}
}

func TestIsHolidayAlt(t *testing.T) {
	var t1 time.Time

	t1 = ezTime(2022, 7, 22)
	t.Run("Darwin Show Day Check 2022", testIsHolidayAlt(NT, t1, true))

	t1 = ezTime(2023, 7, 28)
	t.Run("Darwin Show Day Check 2023", testIsHolidayAlt(NT, t1, true))

	t1 = ezTime(2024, 7, 26)
	t.Run("Darwin Show Day Check 2024", testIsHolidayAlt(NT, t1, true))

	t1 = ezTime(2025, 7, 25)
	t.Run("Darwin Show Day Check 2025", testIsHolidayAlt(NT, t1, true))

	t1 = ezTime(2026, 7, 24)
	t.Run("Darwin Show Day Check 2026", testIsHolidayAlt(NT, t1, true))

	t1 = ezTime(2026, 7, 25)
	t.Run("Darwin Show Day Check False", testIsHolidayAlt(NT, t1, false))
}

func testIsHolidayAlt(state State, date time.Time, expect bool) func(t *testing.T) {
	return func(t *testing.T) {
		b := IsHolidayAlt(state, date)
		if b != expect {
			t.Errorf("IsHoliday was incorrect, got: %t, want: %t", b, expect)
		}
	}

}

func TestBusinessDays(t *testing.T) {
	var t1, t2 time.Time

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 17)
	t.Run("2019-04-17,17", testBusinessDaysFunc(ausHDay, t1, t2, 0))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 10)
	t.Run("2019-04-17,10", testBusinessDaysFunc(ausHDay, t1, t2, -1))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 18)
	t.Run("2019-04-17,18", testBusinessDaysFunc(ausHDay, t1, t2, 1))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 22)
	t.Run("2019-04-17,22", testBusinessDaysFunc(ausHDay, t1, t2, 1))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 24)
	t.Run("2019-04-17,24", testBusinessDaysFunc(ausHDay, t1, t2, 3))
}
func testBusinessDaysFunc(hday *AusHoliday, t1, t2 time.Time, expect int) func(t *testing.T) {
	return func(t *testing.T) {
		days := hday.BusinessDaysAny(t1, t2)
		if days != expect {
			t.Errorf("BusinessDays was incorrect, got: %d, want: %d", days, expect)
		}
	}
}

func TestOffDays(t *testing.T) {
	var t1, t2 time.Time

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 17)
	t.Run("2019-04-17,17", testOffDaysFunc(ausHDay, t1, t2, 0))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 10)
	t.Run("2019-04-17,10", testOffDaysFunc(ausHDay, t1, t2, -1))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 18)
	t.Run("2019-04-17,18", testOffDaysFunc(ausHDay, t1, t2, 0))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 22)
	t.Run("2019-04-17,22", testOffDaysFunc(ausHDay, t1, t2, 4))

	t1 = ezTime(2019, 4, 17)
	t2 = ezTime(2019, 4, 24)
	t.Run("2019-04-17,24", testOffDaysFunc(ausHDay, t1, t2, 4))
}
func testOffDaysFunc(hday *AusHoliday, t1, t2 time.Time, expect int) func(t *testing.T) {
	return func(t *testing.T) {
		days := hday.OffDays(ACT, t1, t2)
		if days != expect {
			t.Errorf("Days was incorrect, got: %d, want: %d", days, expect)
		}
	}
}

func TestBusinessDaysState(t *testing.T) {
	t.Run("2019-04-17,24 act", testBusinessDaysStateFunc(ausHDay, "act", d1, d2, 3))
	t.Run("2019-04-17,24 wa", testBusinessDaysStateFunc(ausHDay, "wa", d1, d2, 4))
}
func testBusinessDaysStateFunc(hday *AusHoliday, state string, t1, t2 time.Time, expect int) func(t *testing.T) {
	return func(t *testing.T) {
		days := hday.BusinessDays(States[state], t1, t2)
		if days != expect {
			t.Errorf("BusinessDaysState was incorrect, got: %d, want: %d", days, expect)
		}
	}
}

// ezTime easily create time.Time
func ezTime(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
}

func TestMain(m *testing.M) {
	// specify location
	loc, _ = time.LoadLocation("Australia/Perth")

	// specify time for test
	d1, _ = time.ParseInLocation(time.RFC3339, "2019-04-17T00:00:00+08:00", loc)
	d2, _ = time.ParseInLocation(time.RFC3339, "2019-04-24T00:00:00+08:00", loc)

	// init new holiday with test data
	localCSV := strings.NewReader(csvStrTest)
	ausHDay, _ = NewFromReader(localCSV)

	os.Exit(m.Run())
}
