package ddate

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		d time.Time
	}

	// Some time objects to test against
	var (
		timeFmt          = "2006-01-02"
		nyDay2019, _     = time.Parse(timeFmt, "2019-01-01")
		befStTibs2020, _ = time.Parse(timeFmt, "2020-02-28")
		stTibs2020, _    = time.Parse(timeFmt, "2020-02-29")
		aftStTibs2020, _ = time.Parse(timeFmt, "2020-03-01")
	)

	tests := []struct {
		name string
		args args
		want Date
	}{
		// TODO: Add test cases.
		{
			"New Years Day 2019",
			args{nyDay2019},
			Date{
				Year:        3185,
				Season:      ssnChaos,
				Day:         days[0],
				DayOfWeek:   0,
				DayOfSeason: 1,
			},
		},
		{
			"Day before St Tib's Day 2020",
			args{befStTibs2020},
			Date{
				Year:        3186,
				Season:      ssnChaos,
				Day:         days[3],
				DayOfWeek:   3,
				DayOfSeason: 59,
			},
		},
		{
			"St Tib's Day 2020",
			args{stTibs2020},
			Date{
				Year:        3186,
				Season:      ssnChaos,
				Day:         stTibsDay,
				DayOfWeek:   3,
				DayOfSeason: 59,
			},
		},
		{
			"Day After St Tib's Day 2020",
			args{aftStTibs2020},
			Date{
				Year:        3186,
				Season:      ssnChaos,
				Day:         days[4],
				DayOfWeek:   4,
				DayOfSeason: 60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
