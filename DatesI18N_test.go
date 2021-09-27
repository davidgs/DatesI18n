package datesI18N

import (
	"testing"
)

var testLanguage = "ru"
func Test_NewDatesI18N(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Language != testLanguage {
		t.Error("Error in NewDatesI18N")
	}
}

func Test_DatesI18N_DayName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays.Monday != "" {
		if d.DayName(1) != d.Weekdays.Monday {
			t.Error("Error in DatesI18N_DayName")
		}
	}
}

func Test_DatesI18N_DayNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays.Monday != "" {
		if d.DayNumber(d.Weekdays.Monday) != 1 {
			t.Error("Error in DayNumber")
		}
	}
}

func Test_DatesI18N_ShortDayName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays_Short.Monday != "" {
		if d.ShortDayName(1) != d.Weekdays_Short.Monday {
			t.Error("Error in DatesI18N_ShortDayName")
		}
	}
}

func Test_DatesI18N_ShortDayNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays_Short.Monday != "" {
		if d.ShortDayNumber(d.Weekdays_Short.Monday) != 1 {
			t.Error("Error in ShortDayNumber")
		}
	}
}

func Test_DatesI18N_MinDayName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays_Min.Monday != "" {
		if d.MinDayName(1) != d.Weekdays_Min.Monday {
			t.Error("Error in DatesI18N_ShortDayName")
		}
	}
}

func Test_DatesI18N_MinDayNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Weekdays_Min.Monday != "" {
		if d.MinDayNumber(d.Weekdays_Min.Monday) != 1 {
			t.Error("Error in MinDayNumber")
		}
	}
}

func Test_DatesI18N_MonthName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Months.January != "" {
		if d.MonthName(1) != d.Months.January {
			t.Error("Error in DatesI18N_MonthName")
		}
	}
}

func Test_DatesI18N_ShortMonthName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Months_Short.January != "" {
		if d.ShortMonthName(1) != d.Months_Short.January {
			t.Error("Error in DatesI18N_ShortMonthName")
		}
	}
}

func Test_DatesI18N_ShortMonthNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.Months_Short.January != "" {
		if d.ShortMonthNumber(d.Months_Short.January) != 1 {
			t.Error("Error in ShortMonthNumber")
		}
	}
}
