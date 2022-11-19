package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StringMeeting struct {
	Start string
	End   string
}

type meeting struct {
	Start int
	End   int
}

type calendar []meeting

func CalendarMatching(
	c1 []StringMeeting, dailyBounds1 StringMeeting,
	c2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	return getMatchingVacancies(
		flattenCalendar(
			mergeCalendars(updateCalendar(c1, dailyBounds1), updateCalendar(c2, dailyBounds2)),
		),
		meetingDuration)
}

func updateCalendar(calendar []StringMeeting, dailyBounds StringMeeting) calendar {
	updatedCalendar := append([]StringMeeting{
		{Start: "0:00", End: dailyBounds.Start},
	}, calendar...)
	updatedCalendar = append(updatedCalendar, StringMeeting{
		Start: dailyBounds.End, End: "23:59",
	})

	meetings := []meeting{}
	for _, i := range updatedCalendar {
		meetings = append(meetings, meeting{
			Start: timeToMinutes(i.Start),
			End:   timeToMinutes(i.End),
		})
	}
	return meetings
}

func mergeCalendars(c1, c2 calendar) calendar {
	merged := []meeting{}
	i, j := 0, 0
	for i < len(c1) && j < len(c2) {
		meeting1, meeting2 := c1[i], c2[j]
		if meeting1.Start < meeting2.Start {
			merged = append(merged, meeting1)
			i++
		} else {
			merged = append(merged, meeting2)
			j++
		}
	}

	for i < len(c1) {
		merged = append(merged, c1[i])
		i++
	}
	for j < len(c2) {
		merged = append(merged, c2[j])
		j++
	}
	return merged
}

func flattenCalendar(c calendar) calendar {
	flattened := []meeting{c[0]}
	for i := 1; i < len(c); i++ {
		currentMeeting := c[i]
		previousMeeting := flattened[len(flattened)-1]
		if previousMeeting.End >= currentMeeting.Start {
			newPreviousMeeting := meeting{
				Start: previousMeeting.Start,
				End:   max(previousMeeting.End, currentMeeting.End),
			}
			flattened[len(flattened)-1] = newPreviousMeeting
		} else {
			flattened = append(flattened, currentMeeting)
		}
	}
	return flattened
}

func getMatchingVacancies(c calendar, meetingLen int) []StringMeeting {
	sameVacancy := []StringMeeting{}
	for i := 1; i < len(c); i++ {
		start := c[i-1].End
		end := c[i].Start
		vacancyDuration := end - start
		if vacancyDuration >= meetingLen {
			sameVacancy = append(sameVacancy, StringMeeting{
				Start: minutesToTime(start),
				End:   minutesToTime(end),
			})
		}
	}
	return sameVacancy
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func timeToMinutes(time string) int {
	hour, mins, found := strings.Cut(time, ":")
	if !found {
		return 0
	}
	hours, err := strconv.Atoi(hour)
	if err != nil {
		return 0
	}
	minutes, err := strconv.Atoi(mins)
	if err != nil {
		return 0
	}
	return hours*60 + minutes
}

func minutesToTime(minutes int) string {
	return fmt.Sprintf("%d:%02d", minutes/60, minutes%60)
}

// Test Case 1
//
// {
//   "calendar1": [
//     ["9:00", "10:30"],
//     ["12:00", "13:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:00", "20:00"],
//   "calendar2": [
//     ["10:00", "11:30"],
//     ["12:30", "14:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["10:00", "18:30"],
//   "meetingDuration": 30
// }
//
// Test Case 2
//
// {
//   "calendar1": [
//     ["9:00", "10:30"],
//     ["12:00", "13:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:00", "20:00"],
//   "calendar2": [
//     ["10:00", "11:30"],
//     ["12:30", "14:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["10:00", "18:30"],
//   "meetingDuration": 30
// }
//
// Test Case 3
//
// {
//   "calendar1": [
//     ["9:00", "10:30"],
//     ["12:00", "13:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:00", "20:00"],
//   "calendar2": [
//     ["10:00", "11:30"],
//     ["12:30", "14:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["10:00", "18:30"],
//   "meetingDuration": 45
// }
//
// Test Case 4
//
// {
//   "calendar1": [
//     ["9:00", "10:30"],
//     ["12:00", "13:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["8:00", "20:00"],
//   "calendar2": [
//     ["10:00", "11:30"],
//     ["12:30", "14:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["7:00", "18:30"],
//   "meetingDuration": 45
// }
//
// Test Case 5
//
// {
//   "calendar1": [
//     ["8:00", "10:30"],
//     ["11:30", "13:00"],
//     ["14:00", "16:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["8:00", "18:00"],
//   "calendar2": [
//     ["10:00", "11:30"],
//     ["12:30", "14:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["7:00", "18:30"],
//   "meetingDuration": 15
// }
//
// Test Case 6
//
// {
//   "calendar1": [
//     ["10:00", "10:30"],
//     ["10:45", "11:15"],
//     ["11:30", "13:00"],
//     ["14:15", "16:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:30", "20:00"],
//   "calendar2": [
//     ["10:00", "11:00"],
//     ["12:30", "13:30"],
//     ["14:30", "15:00"],
//     ["16:00", "17:00"]
//   ],
//   "dailyBounds2": ["9:00", "18:30"],
//   "meetingDuration": 15
// }
//
// Test Case 7
//
// {
//   "calendar1": [
//     ["10:00", "10:30"],
//     ["10:45", "11:15"],
//     ["11:30", "13:00"],
//     ["14:15", "16:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:30", "20:00"],
//   "calendar2": [
//     ["10:00", "11:00"],
//     ["10:30", "20:30"]
//   ],
//   "dailyBounds2": ["9:00", "22:30"],
//   "meetingDuration": 60
// }
//
// Test Case 8
//
// {
//   "calendar1": [
//     ["10:00", "10:30"],
//     ["10:45", "11:15"],
//     ["11:30", "13:00"],
//     ["14:15", "16:00"],
//     ["16:00", "18:00"]
//   ],
//   "dailyBounds1": ["9:30", "20:00"],
//   "calendar2": [
//     ["10:00", "11:00"],
//     ["10:30", "16:30"]
//   ],
//   "dailyBounds2": ["9:00", "22:30"],
//   "meetingDuration": 60
// }
//
// Test Case 9
//
// {
//   "calendar1": [],
//   "dailyBounds1": ["9:30", "20:00"],
//   "calendar2": [],
//   "dailyBounds2": ["9:00", "16:30"],
//   "meetingDuration": 60
// }
//
// Test Case 10
//
// {
//   "calendar1": [],
//   "dailyBounds1": ["9:00", "16:30"],
//   "calendar2": [],
//   "dailyBounds2": ["9:30", "20:00"],
//   "meetingDuration": 60
// }
//
// Test Case 11
//
// {
//   "calendar1": [],
//   "dailyBounds1": ["9:30", "16:30"],
//   "calendar2": [],
//   "dailyBounds2": ["9:30", "16:30"],
//   "meetingDuration": 60
// }
//
// Test Case 12
//
// {
//   "calendar1": [
//     ["7:00", "7:45"],
//     ["8:15", "8:30"],
//     ["9:00", "10:30"],
//     ["12:00", "14:00"],
//     ["14:00", "15:00"],
//     ["15:15", "15:30"],
//     ["16:30", "18:30"],
//     ["20:00", "21:00"]
//   ],
//   "dailyBounds1": ["6:30", "22:00"],
//   "calendar2": [
//     ["9:00", "10:00"],
//     ["11:15", "11:30"],
//     ["11:45", "17:00"],
//     ["17:30", "19:00"],
//     ["20:00", "22:15"]
//   ],
//   "dailyBounds2": ["8:00", "22:30"],
//   "meetingDuration": 30
// }
