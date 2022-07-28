package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the day of the week that the WeekSchedule specifies for the
// givin year and month. If there is an unknown problem -1 is returned.
func Day(wSched WeekSchedule, d time.Weekday, m time.Month, y int) int {
	date := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)

	days := findMatchingDays(date, d)
	switch wSched {
	case First:
		return days[0]
	case Second:
		return days[1]
	case Third:
		return days[2]
	case Fourth:
		return days[3]
	case Last:
		return days[len(days)-1]
	case Teenth:
		for _, d := range days {
			if d >= 13 {
				return d
			}
		}
	}
	return -1
}

// findMatchingDays finds all days as ints matching d.
func findMatchingDays(date time.Time, d time.Weekday) []int {
	days := []int{}
	month := date.Month()
	for date.Month() == month {
		if date.Weekday() == d {
			days = append(days, date.Day())
		}
		date = date.Add(24 * time.Hour)
	}
	return days
}
