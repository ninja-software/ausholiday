# AusHoliday

A tool to calculate how many business days between 2 dates.

## Install

go get github.com/ninja-software/ausholiday

## Functions

.Days(...) how many days shared between two date  
.BusinessDays(...) how many business days between two date in state  
.IsWeekend(...) is it a weekend day  
.IsHoliday(...) is it a holiday in State  
.IsHolidayAny(...) is it a holiday in any state  
.BusinessDaysAny(...) how many business days between two date in any state  
.OffDays(...) how many holiday days between two date for state  
.OffDaysAny(...) how many holiday days between two date for any state

## Usage

```go
ausholiday.BusinessDays(ausholiday.State, time.Time, time.Time)
```

## Example

```go
package main

import (
    "fmt"
    "time"

    "github.com/ninja-software/ausholiday"
)

func main() {
    loc, _ := time.LoadLocation("Australia/Perth")

    t1, _ := time.ParseInLocation(time.RFC3339, "2019-04-17T00:00:00+08:00", loc)
    t2, _ := time.ParseInLocation(time.RFC3339, "2019-04-24T00:00:00+08:00", loc)

    bizDays := ausholiday.BusinessDays(ausholiday.WA, t1, t2)

    fmt.Printf("There are %d business days between 17th to 24th of April 2019\n", bizDays)
}
```
