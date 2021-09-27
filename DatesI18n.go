package datesI18N

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var language = "en"

type DatesI18N struct {
	Weekdays       Dates_Weekdays       `json:"weekdays"`
	Weekdays_Short Dates_Weekdays_Short `json:"weekdays_short"`
	Weekdays_Min   Dates_Weekdays_Min   `json:"weekdays_min"`
	Months         Dates_Months         `json:"months"`
	Months_Short   Dates_Months_Short   `json:"months_short"`
	Language 		 string               `json:"language"`
}

type Dates_Weekdays struct {
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
}

type Dates_Weekdays_Short struct {
	Sunday    string `json:"sun"`
	Monday    string `json:"mon"`
	Tuesday   string `json:"tue"`
	Wednesday string `json:"wed"`
	Thursday  string `json:"thu"`
	Friday    string `json:"fri"`
	Saturday  string `json:"sat"`
}

type Dates_Weekdays_Min struct {
	Sunday    string `json:"Su"`
	Monday    string `json:"Mo"`
	Tuesday   string `json:"Tu"`
	Wednesday string `json:"We"`
	Thursday  string `json:"Th"`
	Friday    string `json:"Fr"`
	Saturday  string `json:"Sa"`
}

type Dates_Months struct {
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

type Dates_Months_Short struct {
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

func (d *DatesI18N) ShortDayName(day int) string { // test written
	switch day {
	case 0:
		return d.Weekdays_Short.Sunday
	case 1:
		return d.Weekdays_Short.Monday
	case 2:
		return d.Weekdays_Short.Tuesday
	case 3:
		return d.Weekdays_Short.Wednesday
	case 4:
		return d.Weekdays_Short.Thursday
	case 5:
		return d.Weekdays_Short.Friday
	case 6:
		return d.Weekdays_Short.Saturday
	}
	return ""
}

func (d *DatesI18N) ShortDayNumber(day string) int { // test written
	switch day {
	case d.Weekdays_Short.Sunday:
		return 0
	case d.Weekdays_Short.Monday:
		return 1
	case d.Weekdays_Short.Tuesday:
		return 2
	case d.Weekdays_Short.Wednesday:
		return 3
	case d.Weekdays_Short.Thursday:
		return 4
	case d.Weekdays_Short.Friday:
		return 5
	case d.Weekdays_Short.Saturday:
		return 6
	}
	return -1
}

func (d *DatesI18N) MinDayName(day int) string { // test written
	switch day {
	case 0:
		return d.Weekdays_Min.Sunday
	case 1:
		return d.Weekdays_Min.Monday
	case 2:
		return d.Weekdays_Min.Tuesday
	case 3:
		return d.Weekdays_Min.Wednesday
	case 4:
		return d.Weekdays_Min.Thursday
	case 5:
		return d.Weekdays_Min.Friday
	case 6:
		return d.Weekdays_Min.Saturday
	}
	return ""
}

func (d *DatesI18N) MinDayNumber(day string) int {
	switch day {
	case d.Weekdays_Min.Sunday:
		return 0
	case d.Weekdays_Min.Monday:
		return 1
	case d.Weekdays_Min.Tuesday:
		return 2
	case d.Weekdays_Min.Wednesday:
		return 3
	case d.Weekdays_Min.Thursday:
		return 4
	case d.Weekdays_Min.Friday:
		return 5
	case d.Weekdays_Min.Saturday:
		return 6
	}
	return -1
}

func (d *DatesI18N) MonthName(month int) string {
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

func (d *DatesI18N) MonthNumber(month string) int {
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

func (d *DatesI18N) ShortMonthName(month int) string {
	switch month {
	case 1:
		return d.Months_Short.January
	case 2:
		return d.Months_Short.February
	case 3:
		return d.Months_Short.March
	case 4:
		return d.Months_Short.April
	case 5:
		return d.Months_Short.May
	case 6:
		return d.Months_Short.June
	case 7:
		return d.Months_Short.July
	case 8:
		return d.Months_Short.August
	case 9:
		return d.Months_Short.September
	case 10:
		return d.Months_Short.October
	case 11:
		return d.Months_Short.November
	case 12:
		return d.Months_Short.December
	}
	return ""
}

func (d *DatesI18N) ShortMonthNumber(month string) int {
	switch month {
	case d.Months_Short.January:
		return 1
	case d.Months_Short.February:
		return 2
	case d.Months_Short.March:
		return 3
	case d.Months_Short.April:
		return 4
	case d.Months_Short.May:
		return 5
	case d.Months_Short.June:
		return 6
	case d.Months_Short.July:
		return 7
	case d.Months_Short.August:
		return 8
	case d.Months_Short.September:
		return 9
	case d.Months_Short.October:
		return 10
	case d.Months_Short.November:
		return 11
	case d.Months_Short.December:
		return 12
	}
	return -1
}


func NewDatesI18N(lang string) *DatesI18N { // test written
	d := new(DatesI18N)
	d.setLanguage(lang)
	return d
}

func (d *DatesI18N) setLanguage(lang string) {
	d.Language = lang
	dat, err := ioutil.ReadFile("lang/dates_" + lang + ".json")
	if err != nil {
		fmt.Println("No json file: ", err)
	}
	err = json.Unmarshal(dat, d)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
}