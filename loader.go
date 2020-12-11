package ausholiday

import (
	"encoding/csv"
	"io"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

// NewFromReader initialise from reader stream
// use for custom holiday dates
func NewFromReader(csvfile io.Reader) (*AusHoliday, error) {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)

		r.Comma = ','
		return r // Allows use customer character as delimiter
	})

	rows := []*inputCSV{}
	err := gocsv.Unmarshal(csvfile, &rows)
	if err != nil {
		return nil, err
	}

	ymd := "20060102"
	loc, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		return nil, err
	}

	holidays := []*Holiday{}
	for _, row := range rows {
		// work out time using unix time
		t, err := time.ParseInLocation(ymd, row.Date, loc)
		if err != nil {
			return nil, err
		}

		// note this will basically accept any text, so it doesnt do sanity check on state
		// TODO add sanity check
		jxn := States[strings.ToLower(row.Jurisdiction)]

		holidays = append(holidays, &Holiday{
			Date:            t,
			YMD:             row.Date,
			HolidayName:     row.HolidayName,
			Information:     row.Information,
			MoreInformation: row.MoreInformation,
			Jurisdiction:    jxn,
		})
	}

	hDay := &AusHoliday{
		Holidays: holidays,
	}

	return hDay, nil
}

// Days how many days shared between two date
func (ahd *AusHoliday) Days(from, to time.Time) int {
	return Days(from, to)
}

// IsWeekend is it a weekend day
func (ahd *AusHoliday) IsWeekend(t time.Time) bool {
	return IsWeekend(t)
}

// BusinessDays how many business days between two date
func (ahd *AusHoliday) BusinessDays(state State, from, to time.Time) int {
	return BusinessDays(state, from, to, ahd.Holidays...)
}

// BusinessDaysAny how many business days between two date in state in any state
func (ahd *AusHoliday) BusinessDaysAny(from, to time.Time) int {
	return BusinessDaysAny(from, to, ahd.Holidays...)
}

// OffDays how many holiday days between two date for state
func (ahd *AusHoliday) OffDays(state State, from, to time.Time) int {
	return OffDays(state, from, to, ahd.Holidays...)
}

// OffDaysAny how many holiday days between two date in any state
func (ahd *AusHoliday) OffDaysAny(from, to time.Time) int {
	return OffDaysAny(from, to, ahd.Holidays...)
}

// IsHoliday is it a holiday in State
func (ahd *AusHoliday) IsHoliday(state State, t time.Time) bool {
	return IsHoliday(state, t, ahd.Holidays...)
}

// IsHolidayAny is it a holiday in any state
func (ahd *AusHoliday) IsHolidayAny(t time.Time) bool {
	return IsHolidayAny(t, ahd.Holidays...)
}
