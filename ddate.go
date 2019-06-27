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
	seasonID, dayOfSeason := "", 0
	for season, days := range seasons {
		if yearDay >= days.first && yearDay <= days.last {
			seasonID = season
			dayOfSeason = (yearDay - days.first) + 1
			break
		}
	}

	return Date{
		Year:        (d.Year() - 1970) + 3136,
		Season:      seasonID,
		DayOfSeason: dayOfSeason,
		DayOfWeek:   dayOfSeason % 5,
		Day:         days[dayOfSeason%5],
	}
}

// Today returns Today's date in the discordian calendar
func Today() Date {
	return New(time.Now())
}

// Internals used for calculating things
type dateRange struct {
	first int
	last  int
}

var seasons = map[string]dateRange{
	"Chaos":         {1, 73},
	"Discord":       {74, 146},
	"Confusion":     {147, 219},
	"Bureaucracy":   {220, 292},
	"The Aftermath": {293, 365},
}

var days = []string{"Sweetmorn", "Boomtime", "Pungenday", "Prickle-Prickle", "Setting Orange"}
