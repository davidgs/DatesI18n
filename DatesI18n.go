// MIT License

// Copyright (c) 2021 David G. Simmons

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package datesi18n provides internationalization for programs needing
// language support for weekdays and months. The package is intended to
// be used with the dates package.
//
// The following code is an example of using the datesi18n package:
//
//	import "github.com/mjibson/go-datesi18n"
//
//	func main() {
//		d := datesi18n.NewDatesI18N("en")
//		fmt.Println(d.Weekdays.Monday)
//		fmt.Println(d.WeekdaysShort.Monday)
//		fmt.Println(d.WeekdaysMin.Monday)
//		fmt.Println(d.Months.January)
//	}
//
// The output of the program is:
//
//	Monday
//	Mon
//	M
//	January
package datesi18n

import (
	"embed"
	"encoding/json"
	"fmt"
)

var language = "en"

// DatesI18N is a struct that holds the internationalization data for
// the dates package.
type DatesI18N struct {
	Weekdays      DatesWeekdays      `json:"weekdays"`
	WeekdaysShort DatesWeekdaysShort `json:"weekdaysshort"`
	WeekdaysMin   DatesWeekdaysMin   `json:"weekdaysmin"`
	Months        DatesMonths        `json:"months"`
	MonthsShort   DatesMonthsShort   `json:"monthsshort"`
	Language      string             `json:"language"`
}

//go:embed lang/*.json
var f embed.FS

// DatesWeekdays is a struct that holds the names of the weekdays.
type DatesWeekdays struct {
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
}

// DatesWeekdaysShort is a struct that holds the short names of the weekdays.
type DatesWeekdaysShort struct {
	Sunday    string `json:"sun"`
	Monday    string `json:"mon"`
	Tuesday   string `json:"tue"`
	Wednesday string `json:"wed"`
	Thursday  string `json:"thu"`
	Friday    string `json:"fri"`
	Saturday  string `json:"sat"`
}

// DattesWeekdaysMin is a struct that holds the minimal names of the weekdays.
type DatesWeekdaysMin struct {
	Sunday    string `json:"Su"`
	Monday    string `json:"Mo"`
	Tuesday   string `json:"Tu"`
	Wednesday string `json:"We"`
	Thursday  string `json:"Th"`
	Friday    string `json:"Fr"`
	Saturday  string `json:"Sa"`
}

// DatesMonths is a struct that holds the names of the months.
type DatesMonths struct {
	January   string `json:"january"`
	February  string `json:"february"`
	March     string `json:"march"`
	April     string `json:"april"`
	May       string `json:"may"`
	June      string `json:"june"`
	July      string `json:"july"`
	August    string `json:"august"`
	September string `json:"september"`
	October   string `json:"october"`
	November  string `json:"november"`
	December  string `json:"december"`
}

// DatesMonthsShort is a struct that holds the short names of the months.
type DatesMonthsShort struct {
	January   string `json:"jan"`
	February  string `json:"feb"`
	March     string `json:"mar"`
	April     string `json:"apr"`
	May       string `json:"may"`
	June      string `json:"jun"`
	July      string `json:"jul"`
	August    string `json:"aug"`
	September string `json:"sep"`
	October   string `json:"oct"`
	November  string `json:"nov"`
	December  string `json:"dec"`
}

// DayName returns the name of the day of the week given by index number.
func (d *DatesI18N) DayName(day int) string { // test written
	switch day {
	case 0:
		return d.Weekdays.Sunday
	case 1:
		return d.Weekdays.Monday
	case 2:
		return d.Weekdays.Tuesday
	case 3:
		return d.Weekdays.Wednesday
	case 4:
		return d.Weekdays.Thursday
	case 5:
		return d.Weekdays.Friday
	case 6:
		return d.Weekdays.Saturday
	}
	return ""
}

// DayNumber returns the number of the day of the week given by name.
func (d *DatesI18N) DayNumber(day string) int { // test written
	switch day {
	case d.Weekdays.Sunday:
		return 0
	case d.Weekdays.Monday:
		return 1
	case d.Weekdays.Tuesday:
		return 2
	case d.Weekdays.Wednesday:
		return 3
	case d.Weekdays.Thursday:
		return 4
	case d.Weekdays.Friday:
		return 5
	case d.Weekdays.Saturday:
		return 6
	}
	return -1
}

// ShortDayName returns the short name of the day of the week given by index number.
func (d *DatesI18N) ShortDayName(day int) string { // test written
	switch day {
	case 0:
		return d.WeekdaysShort.Sunday
	case 1:
		return d.WeekdaysShort.Monday
	case 2:
		return d.WeekdaysShort.Tuesday
	case 3:
		return d.WeekdaysShort.Wednesday
	case 4:
		return d.WeekdaysShort.Thursday
	case 5:
		return d.WeekdaysShort.Friday
	case 6:
		return d.WeekdaysShort.Saturday
	}
	return ""
}

// ShortDayNumber returns the number of the day of the week given by short name.
func (d *DatesI18N) ShortDayNumber(day string) int { // test written
	switch day {
	case d.WeekdaysShort.Sunday:
		return 0
	case d.WeekdaysShort.Monday:
		return 1
	case d.WeekdaysShort.Tuesday:
		return 2
	case d.WeekdaysShort.Wednesday:
		return 3
	case d.WeekdaysShort.Thursday:
		return 4
	case d.WeekdaysShort.Friday:
		return 5
	case d.WeekdaysShort.Saturday:
		return 6
	}
	return -1
}

// MinDayName returns the minimal name of the day of the week given by index number.
func (d *DatesI18N) MinDayName(day int) string { // test written
	switch day {
	case 0:
		return d.WeekdaysMin.Sunday
	case 1:
		return d.WeekdaysMin.Monday
	case 2:
		return d.WeekdaysMin.Tuesday
	case 3:
		return d.WeekdaysMin.Wednesday
	case 4:
		return d.WeekdaysMin.Thursday
	case 5:
		return d.WeekdaysMin.Friday
	case 6:
		return d.WeekdaysMin.Saturday
	}
	return ""
}

// MinDayNumber returns the number of the day of the week given by minimal name.
func (d *DatesI18N) MinDayNumber(day string) int {
	switch day {
	case d.WeekdaysMin.Sunday:
		return 0
	case d.WeekdaysMin.Monday:
		return 1
	case d.WeekdaysMin.Tuesday:
		return 2
	case d.WeekdaysMin.Wednesday:
		return 3
	case d.WeekdaysMin.Thursday:
		return 4
	case d.WeekdaysMin.Friday:
		return 5
	case d.WeekdaysMin.Saturday:
		return 6
	}
	return -1
}

// MonthName returns the name of the month given by index number.
func (d *DatesI18N) MonthName(month int) string {
	month += iota
	switch month {
	case 1:
		return d.Months.January
	case 2:
		return d.Months.February
	case 3:
		return d.Months.March
	case 4:
		return d.Months.April
	case 5:
		return d.Months.May
	case 6:
		return d.Months.June
	case 7:
		return d.Months.July
	case 8:
		return d.Months.August
	case 9:
		return d.Months.September
	case 10:
		return d.Months.October
	case 11:
		return d.Months.November
	case 12:
		return d.Months.December
	}
	return ""
}

// MonthNumber returns the number of the month given by name.
func (d *DatesI18N) MonthNumber(month string) int {
	month += iota
	switch month {
	case d.Months.January:
		return 1
	case d.Months.February:
		return 2
	case d.Months.March:
		return 3
	case d.Months.April:
		return 4
	case d.Months.May:
		return 5
	case d.Months.June:
		return 6
	case d.Months.July:
		return 7
	case d.Months.August:
		return 8
	case d.Months.September:
		return 9
	case d.Months.October:
		return 10
	case d.Months.November:
		return 11
	case d.Months.December:
		return 12
	}
	return -1
}

// ShortMonthName returns the short name of the month given by index number.
func (d *DatesI18N) ShortMonthName(month int) string {
	month += iota
	switch month {
	case 1:
		return d.MonthsShort.January
	case 2:
		return d.MonthsShort.February
	case 3:
		return d.MonthsShort.March
	case 4:
		return d.MonthsShort.April
	case 5:
		return d.MonthsShort.May
	case 6:
		return d.MonthsShort.June
	case 7:
		return d.MonthsShort.July
	case 8:
		return d.MonthsShort.August
	case 9:
		return d.MonthsShort.September
	case 10:
		return d.MonthsShort.October
	case 11:
		return d.MonthsShort.November
	case 12:
		return d.MonthsShort.December
	}
	return ""
}

// ShortMonthNumber returns the number of the month given by short name.
func (d *DatesI18N) ShortMonthNumber(month string) int {
	switch month {
	case d.MonthsShort.January:
		return 1
	case d.MonthsShort.February:
		return 2
	case d.MonthsShort.March:
		return 3
	case d.MonthsShort.April:
		return 4
	case d.MonthsShort.May:
		return 5
	case d.MonthsShort.June:
		return 6
	case d.MonthsShort.July:
		return 7
	case d.MonthsShort.August:
		return 8
	case d.MonthsShort.September:
		return 9
	case d.MonthsShort.October:
		return 10
	case d.MonthsShort.November:
		return 11
	case d.MonthsShort.December:
		return 12
	}
	return -1
}

// NewDatesI18N returns a new DatesI18N struct initialized with the language values for the given language
func NewDatesI18N(lang string) *DatesI18N { // test written
	d := new(DatesI18N)
	d.setLanguage(lang)
	return d
}

// loads the language file for the given language
func (d *DatesI18N) setLanguage(lang string) {
	d.Language = lang
	dat, err := f.ReadFile("lang/dates_"+lang+".json")
	if err != nil {
		fmt.Println("No json file: ", err)
	}
	err = json.Unmarshal(dat, d)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
}
