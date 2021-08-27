package dateparse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseAny(t *testing.T) {
	diff := []struct {
		TestName  string
		SMEsDates string
		PrepDates string
		DiffDates []time.Time
	}{

		{
			TestName:  "Simple/Dot/American",
			SMEsDates: "09.Apr.2020",

			PrepDates: "2020-04-09",
		},
		{
			TestName:  "Simple/Dot/American/full",
			SMEsDates: "01.June.2020",

			PrepDates: "2020-06-01",
		},

		{
			TestName:  "Simple/Dash/American",
			SMEsDates: "27-May-20",

			PrepDates: "2020-05-27",
		},

		{
			TestName:  "Simple/Dot/American/CAPS",
			SMEsDates: "25.FEB.2021",

			PrepDates: "2021-02-25",
		},
		{
			TestName:  "Simple/Space/American",
			SMEsDates: "27 May 2020 ",

			PrepDates: "2020-05-27",
		},
		{
			TestName:  "Simple/Slash",
			SMEsDates: "04/14/2020 ",

			PrepDates: "2020-04-14",
		},
		{
			TestName:  "Slash/WordMounth",
			SMEsDates: "06/Apr/20",

			PrepDates: "2020-04-06",
		},
		{
			TestName:  "Slash/WordMounth/Full",
			SMEsDates: "06/Apr/2020",

			PrepDates: "2020-04-06",
		},
		{
			TestName:  "Numerate/Dot/Russian",
			SMEsDates: "03.апр.2020",

			PrepDates: "2020-04-03",
		},
		{
			TestName:  "Numerate/Dot",
			SMEsDates: " 30.11.2020 ",

			PrepDates: "2020-11-30",
		},
		{
			TestName:  "Numerate/American",
			SMEsDates: " 08-APR-2020 ",

			PrepDates: "2020-04-08",
		},
		{
			TestName: "Numerate/Dot/American",

			SMEsDates: " 14.apr.2020 ",
			PrepDates: "2020-04-14",
		},

		{
			TestName: "Numerate/Slash",

			SMEsDates: " 03/03/2021 ",
			PrepDates: "2021-03-03",
		},
		{
			TestName: "Numerate/Dash/American",

			SMEsDates: "  17-APR-2020 ",
			PrepDates: "2020-04-17",
		},
		{
			TestName: "Numerate/Dash/American/CAPS",

			SMEsDates: " 14-MAY-2020 ",
			PrepDates: "2020-05-14",
		},
		{
			TestName: "Numerate/Dot/American/CAPS",

			SMEsDates: " 11.JUN.2020 ",
			PrepDates: "2020-06-11",
		},
		{
			TestName: "Numerate/Space/American",

			SMEsDates: " 09 Jul 2020 ",
			PrepDates: "2020-07-09",
		},
	}

	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {

			dt, err := time.Parse(layoutISO, tt.PrepDates)
			if err != nil {
				t.Error(err)
			}

			SMEDate, err := ParseAny(tt.SMEsDates)
			if err != nil {

				t.Error(err)
				return
			}

			assert.Equal(t, dt, SMEDate)
		})
	}
}
func TestGetMonth(t *testing.T) {
	diff := []struct {
		TestName string
		MonthS   string
		ResS     string
	}{
		//NUMBER
		//____Simple_without_numerate
		{
			TestName: "American/Short",
			MonthS:   "Apr",
			ResS:     "04",
		},
		{
			TestName: "American/full",
			MonthS:   "June",
			ResS:     "06",
		},

		{
			TestName: "Simple/Dash/American",
			MonthS:   "May",
			ResS:     "05",
		},

		{
			TestName: "Simple/Dot/American/CAPS",
			MonthS:   "FEB",
			ResS:     "02",
		},
		{
			TestName: "Simple/Space/American",
			MonthS:   "May",
			ResS:     "05",
		},
		{
			TestName: "Simple/Slash",
			MonthS:   "04",
			ResS:     "04",
		},

		{
			TestName: "Numerate/Dot/Russian",
			MonthS:   "апр",
			ResS:     "04",
		},
		{
			TestName: "Numerate/Dot",
			MonthS:   "11",
			ResS:     "11",
		},
		{
			TestName: "Numerate/American",
			MonthS:   "APR",
			ResS:     "04",
		},
		{
			TestName: "Numerate/Dot/American",
			MonthS:   "apr",
			ResS:     "04",
		},

		{
			TestName: "Numerate/Slash",
			MonthS:   "03",
			ResS:     "03",
		},
		{
			TestName: "Numerate/Dash/American",
			MonthS:   "APR",
			ResS:     "04",
		},
		{
			TestName: "Numerate/Dash/American/CAPS",
			MonthS:   "MAY",
			ResS:     "05",
		},
		{
			TestName: "Numerate/Dot/American/CAPS",
			MonthS:   "JUN",
			ResS:     "06",
		},
		{
			TestName: "Numerate/Space/American",
			MonthS:   "Jul",
			ResS:     "07",
		},
	}

	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {

			month, err := GetMonth(tt.MonthS)
			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, tt.ResS, month)
		})
	}
}
