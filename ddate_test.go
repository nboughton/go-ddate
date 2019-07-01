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
		mid2019, _       = time.Parse(timeFmt, "2019-06-28")
		xmas2019, _      = time.Parse(timeFmt, "2019-12-25")
		befStTibs2020, _ = time.Parse(timeFmt, "2020-02-28")
		stTibs2020, _    = time.Parse(timeFmt, "2020-02-29")
		aftStTibs2020, _ = time.Parse(timeFmt, "2020-03-01")
		tenthJan2019, _  = time.Parse(timeFmt, "2019-01-10")
	)

	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			"New Years Day 2019",
			args{nyDay2019},
			Date{
				year:        3185,
				season:      Chaos,
				day:         days[1],
				dayOfWeek:   1,
				dayOfSeason: 1,
			},
		},
		{
			"Jan 10th 2019",
			args{tenthJan2019},
			Date{
				year:        3185,
				season:      Chaos,
				day:         days[5],
				dayOfWeek:   5,
				dayOfSeason: 10,
			},
		},
		{
			"Mid 2019",
			args{mid2019},
			Date{
				year:        3185,
				season:      Confusion,
				day:         days[4],
				dayOfWeek:   4,
				dayOfSeason: 33,
			},
		},
		{
			"Xmas 2019",
			args{xmas2019},
			Date{
				year:        3185,
				season:      Aftermath,
				day:         days[4],
				dayOfWeek:   4,
				dayOfSeason: 67,
			},
		},
		{
			"Day before St Tib's Day 2020",
			args{befStTibs2020},
			Date{
				year:        3186,
				season:      Chaos,
				day:         days[4],
				dayOfWeek:   4,
				dayOfSeason: 59,
			},
		},
		{
			"St Tib's Day 2020",
			args{stTibs2020},
			Date{
				year:        3186,
				season:      Chaos,
				day:         stTibsDay,
				dayOfWeek:   0,
				dayOfSeason: 0,
			},
		},
		{
			"Day After St Tib's Day 2020",
			args{aftStTibs2020},
			Date{
				year:        3186,
				season:      Chaos,
				day:         days[5],
				dayOfWeek:   5,
				dayOfSeason: 60,
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

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1904", args{1904}, true},
		{"1908", args{1908}, true},
		{"1912", args{1912}, true},
		{"1916", args{1916}, true},
		{"1920", args{1920}, true},
		{"1924", args{1924}, true},
		{"1928", args{1928}, true},
		{"1932", args{1932}, true},
		{"1936", args{1936}, true},
		{"1940", args{1940}, true},
		{"1944", args{1944}, true},
		{"1948", args{1948}, true},
		{"1952", args{1952}, true},
		{"1956", args{1956}, true},
		{"1960", args{1960}, true},
		{"1964", args{1964}, true},
		{"1968", args{1968}, true},
		{"1972", args{1972}, true},
		{"1976", args{1976}, true},
		{"1980", args{1980}, true},
		{"1984", args{1984}, true},
		{"1988", args{1988}, true},
		{"1992", args{1992}, true},
		{"1996", args{1996}, true},
		{"2004", args{2004}, true},
		{"2008", args{2008}, true},
		{"2012", args{2012}, true},
		{"2016", args{2016}, true},
		{"2017", args{2017}, false},
		{"2018", args{2018}, false},
		{"2019", args{2019}, false},
		{"2020", args{2020}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
