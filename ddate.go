package ddate

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// Date represents all the data relevant to a Discordian calender entry
type Date struct {
	Year        int
	Season      string
	Day         string
	DayOfWeek   int
	DayOfSeason int
}

func (d Date) String() string {
	return fmt.Sprintf("Today is %s, %s day of %s in the YOLD %d", d.Day, humanize.Ordinal(d.DayOfSeason), d.Season, d.Year)
}

// New generates a Date from d
func New(d time.Time) Date {
	yearDay := d.YearDay()

	// get the day number and season
	season, dayOfSeason, day := "", 0, ""
	for name, days := range seasons {
		if yearDay >= days.first && yearDay <= days.last {
			season = name
			dayOfSeason = (yearDay - days.first) + 1
			break
		}
	}

	switch dayOfSeason {
	case 5:
		day = seasons[season].apostleDay
	case 50:
		day = seasons[season].seasonDay
	default:
		day = days[dayOfSeason%5]
	}

	return Date{
		Year:        (d.Year() - 1970) + 3136,
		Season:      season,
		DayOfSeason: dayOfSeason,
		DayOfWeek:   dayOfSeason % 5,
		Day:         day,
	}
}

// Today returns Today's date in the discordian calendar
func Today() Date {
	return New(time.Now())
}

// Internals used for calculating things
type seasonData struct {
	first      int
	last       int
	apostleDay string
	seasonDay  string
}

var seasons = map[string]seasonData{
	"Chaos":         {1, 73, "MungDay", "Chaoflux"},
	"Discord":       {74, 146, "Mojoday", "Discoflux"},
	"Confusion":     {147, 219, "Syaday", "Confuflux"},
	"Bureaucracy":   {220, 292, "Zaraday", "Bureflux"},
	"The Aftermath": {293, 365, "Maladay", "Afflux"},
}

var days = []string{"Sweetmorn", "Boomtime", "Pungenday", "Prickle-Prickle", "Setting Orange"}
