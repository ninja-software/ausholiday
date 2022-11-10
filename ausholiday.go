// Copyright 2019 Ninja Software. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

// Package ausholiday give easy access to check for Australian holiday
//
// Can be used immediately or create as a instance with custom holiday dates. Both singleton and instance ausholiday have same functions.
package ausholiday

import (
	"strings"
	"sync"
	"time"
)

// Version of release for ausholiday
const Version = "v1.1.2"

var onceAusHoliday sync.Once
var xholidays []*Holiday

// AusHoliday struct for australia holiday
type AusHoliday struct {
	Holidays []*Holiday
}

// Holiday date
type Holiday struct {
	Date            time.Time
	YMD             string // yyyymmdd
	HolidayName     string
	Information     string
	MoreInformation string
	Jurisdiction    State
}

// State of Australia
type State string

// State data
// https://en.wikipedia.org/wiki/States_and_territories_of_Australia
// order: abbrev, iso
const (
	WA   State = "wa"
	NSW  State = "nsw"
	QLD  State = "qld"
	SA   State = "sa"
	TAS  State = "tas"
	VIC  State = "vic"
	NT   State = "nt"
	ACT  State = "act"
	JBT  State = "jbt"
	AAT  State = "aat"
	CT   State = "ct"   // ISO Christmas Island
	CC   State = "cc"   // ISO Cocos (Keeling) Islands
	CSI  State = "csi"  // Coral Sea Islands
	HIMI State = "himi" // Heard Island and McDonald Islands
	NF   State = "nf"   // ISO NF
)

// States string map of states for quick matching
var States = map[string]State{
	"wa":   WA,
	"nsw":  NSW,
	"qld":  QLD,
	"sa":   SA,
	"tas":  TAS,
	"vic":  VIC,
	"nt":   NT,
	"act":  ACT,
	"jbt":  JBT,
	"aat":  AAT,
	"ct":   CT,
	"cc":   CC,
	"csi":  CSI,
	"himi": HIMI,
	"nf":   NF,
}

// inputCSV csv file structure
type inputCSV struct {
	Date            string `csv:"Date"`
	HolidayName     string `csv:"Holiday Name"`
	Information     string `csv:"Information"`
	MoreInformation string `csv:"More Information"`
	Jurisdiction    string `csv:"Jurisdiction"`
}

func init() {
	// load as singleton, loads built-in holidays
	onceAusHoliday.Do(func() {
		localCSV := strings.NewReader(csvStr)

		days, err := NewFromReader(localCSV)
		if err != nil {
			panic(err)
		}

		xholidays = days.Holidays
	})
}

// Holiday alternative
type HolidayAlt struct {
	HolidayName     string
	Information     string
	MoreInformation string
	Jurisdiction    State
	Timezone        string

	Month   time.Month
	Week    int
	Weekday time.Weekday
}

var holidayAlts = []*HolidayAlt{
	{
		HolidayName:     "Dawin Show Day",
		Information:     "",
		MoreInformation: "",
		Jurisdiction:    NT,
		Timezone:        "Australia/Darwin",

		Month:   time.July,
		Week:    4,
		Weekday: time.Friday,
	},
}

func IsHolidayAlt(state State, date time.Time) bool {

	for _, holidayAlt := range holidayAlts {
		if holidayAlt.Jurisdiction != state {
			continue
		}

		location, err := time.LoadLocation(holidayAlt.Timezone)
		if err != nil {
			continue
		}

		// convert time in the timezone
		timeWithTimezone := date.In(location)

		y := timeWithTimezone.Year()
		m := timeWithTimezone.Month()
		d := timeWithTimezone.Day()
		wd := timeWithTimezone.Weekday()

		// check month
		if holidayAlt.Month != m {
			continue
		}

		// check weekday
		if holidayAlt.Weekday != wd {
			continue
		}

		// check date
		firstDayOfTheMonth := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)

		firstWeekdayOfTheMonth := int(firstDayOfTheMonth.Weekday())

		diffOfTheDays := int(wd) - firstWeekdayOfTheMonth
		if diffOfTheDays < 0 {
			diffOfTheDays += 7
		}

		if d != firstDayOfTheMonth.Add(time.Duration(diffOfTheDays*24)*time.Hour).Add(time.Duration((holidayAlt.Week-1)*7*24)*time.Hour).Day() {
			continue
		}

		return true

	}

	return false
}
