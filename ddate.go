package ddate

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

// Constants, but is anything truly constant?
const (
	stTibsDay   = "St Tib's Day"
	Chaos       = "Chaos"
	Discord     = "Discord"
	Confusion   = "Confusion"
	Bureaucracy = "Bureaucracy"
	Aftermath   = "The Aftermath"
)

// What's in a name?
type attributes struct {
	first      int
	last       int
	apostleDay string
	seasonDay  string
}

var (
	seasons = map[string]attributes{
		Chaos:       {1, 73, "MungDay", "Chaoflux"},
		Discord:     {74, 146, "Mojoday", "Discoflux"},
		Confusion:   {147, 219, "Syaday", "Confuflux"},
		Bureaucracy: {220, 292, "Zaraday", "Bureflux"},
		Aftermath:   {293, 365, "Maladay", "Afflux"},
	}

	days = map[int]string{
		1: "Sweetmorn",
		2: "Boomtime",
		3: "Pungenday",
		4: "Prickle-Prickle",
		5: "Setting Orange",
	}
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

// Date represents a date (SHOCKING!)
type Date struct {
	year        int
	season      string
	day         string
	dayOfWeek   int
	dayOfSeason int
}

// Day returns the name of the current day
func (d Date) Day() string {
	return d.day
}

// DayOfWeek returns the number of the day of the week (1 - 5)
func (d Date) DayOfWeek() int {
	return d.dayOfWeek
}

// DayOfSeason returns the number of the day of the season (1 - 73)
func (d Date) DayOfSeason() int {
	return d.dayOfSeason
}

// Season returns the name of the Season
func (d Date) Season() string {
	return d.season
}

// Year returns the YOLD
func (d Date) Year() int {
	return d.year
}

// String implements the Stringer interface for Date structs
func (d Date) String() string {
	if d.day == stTibsDay {
		return fmt.Sprintf("%s, YOLD %d", d.day, d.year)
	}

	return fmt.Sprintf("%s, %s day of %s in the YOLD %d", d.day, humanize.Ordinal(d.dayOfSeason), d.season, d.year)
}

// New generates a discordian Date from d
func New(d time.Time) Date {
	yearDay := d.YearDay()

	// There's an extra day between 59th and 60th of Chaos on Leap years
	if isLeapYear(d.Year()) && yearDay >= 60 {
		yearDay--
	}

	// Good lord this is all a bit fancy
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

	// We got Holy for Days yo
	switch dayOfSeason {
	case 5:
		day = seasons[season].apostleDay
	case 50:
		day = seasons[season].seasonDay
	default:
		day = days[dayOfWeek]
	}

	// it's St Tibs Day on Feb 29th
	if d.Month() == 2 && d.Day() == 29 {
		day = stTibsDay
		// St Tibs is very much outside of the normal rules of time and space
		dayOfSeason = 0
		dayOfWeek = 0
	}

	return Date{
		year:        (d.Year() - 1970) + 3136, // "What year is it?!"
		season:      season,
		dayOfSeason: dayOfSeason,
		dayOfWeek:   dayOfWeek,
		day:         day,
	}
}

// Today returns Today
func Today() Date {
	return New(time.Now())
}
