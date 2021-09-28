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

package datesi18n

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
	if d.WeekdaysShort.Monday != "" {
		if d.ShortDayName(1) != d.WeekdaysShort.Monday {
			t.Error("Error in DatesI18N_ShortDayName")
		}
	}
}

func Test_DatesI18N_ShortDayNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.WeekdaysShort.Monday != "" {
		if d.ShortDayNumber(d.WeekdaysShort.Monday) != 1 {
			t.Error("Error in ShortDayNumber")
		}
	}
}

func Test_DatesI18N_MinDayName(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.WeekdaysMin.Monday != "" {
		if d.MinDayName(1) != d.WeekdaysMin.Monday {
			t.Error("Error in DatesI18N_ShortDayName")
		}
	}
}

func Test_DatesI18N_MinDayNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.WeekdaysMin.Monday != "" {
		if d.MinDayNumber(d.WeekdaysMin.Monday) != 1 {
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
	if d.MonthsShort.January != "" {
		if d.ShortMonthName(1) != d.MonthsShort.January {
			t.Error("Error in DatesI18N_ShortMonthName")
		}
	}
}

func Test_DatesI18N_ShortMonthNumber(t *testing.T) {
	d := NewDatesI18N(testLanguage)
	if d.MonthsShort.January != "" {
		if d.ShortMonthNumber(d.MonthsShort.January) != 1 {
			t.Error("Error in ShortMonthNumber")
		}
	}
}
