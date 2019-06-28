package ddate

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// Date represents all the data relevant to a Discordian calender entry, note that DayOfWeek is 0 indexed (i.e
// Sweetmorn is 0, Boomtime is 1 etc)
type Date struct {
	Year        int
	Season      string
	Day         string
	DayOfWeek   int
	DayOfSeason int
}

func (d Date) String() string {
	if d.Day == stTibsDay {
		return fmt.Sprintf("Today is %s, YOLD %d", d.Day, d.Year)
	}

	return fmt.Sprintf("Today is %s, %s day of %s in the YOLD %d", d.Day, humanize.Ordinal(d.DayOfSeason), d.Season, d.Year)
}

// New generates a Date from d
func New(d time.Time) Date {
	yearDay := d.YearDay()

	// There's an extra day added between 58th and 59th of Chaos
	if isLeapYear(d.Year()) && yearDay >= 60 {
		yearDay--
	}

	// get the day number and season
	season, day, dayOfSeason, dayOfWeek := "", "", 0, 0
	for name, days := range seasons {
		if yearDay >= days.first && yearDay <= days.last {
			season = name
			dayOfSeason = (yearDay - days.first) + 1

			dayOfWeek = yearDay % 5
			if dayOfWeek == 0 {
				dayOfWeek = 5
			}

			break
		}
	}

	switch dayOfSeason {
	case 5:
		day = seasons[season].apostleDay
	case 50:
		day = seasons[season].seasonDay
	default:
		day = days[dayOfWeek]
	}

	// Implement St Tibbs Day on Feb 29th
	if d.Month() == 2 && d.Day() == 29 {
		day = stTibsDay
		dayOfSeason = 0
		dayOfWeek = 0
	}

	return Date{
		Year:        (d.Year() - 1970) + 3136,
		Season:      season,
		DayOfSeason: dayOfSeason,
		DayOfWeek:   dayOfWeek,
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

var (
	seasons = map[string]seasonData{
		ssnChaos:       {1, 73, "MungDay", "Chaoflux"},
		ssnDiscord:     {74, 146, "Mojoday", "Discoflux"},
		ssnConfusion:   {147, 219, "Syaday", "Confuflux"},
		ssnBureaucracy: {220, 292, "Zaraday", "Bureflux"},
		ssnAftermath:   {293, 365, "Maladay", "Afflux"},
	}
	days = map[int]string{
		1: "Sweetmorn",
		2: "Boomtime",
		3: "Pungenday",
		4: "Prickle-Prickle",
		5: "Setting Orange",
	}
)

// Constant names
const (
	stTibsDay      = "St Tib's Day"
	ssnChaos       = "Chaos"
	ssnDiscord     = "Discord"
	ssnConfusion   = "Confusion"
	ssnBureaucracy = "Bureaucracy"
	ssnAftermath   = "The Aftermath"
)

func isLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
