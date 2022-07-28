package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

// day is a placeholder for the normalize function to normalize hours
const day = 0

// timeRate gives a constant for the x per y, e.g. 24 hours per 1 day
type timeRate int

const (
	mph timeRate = 60 // minutes per hour
	hpd timeRate = 24 // hours per day
)

// New creates a new Clock with provided hours and minutes. The hours and
// minutes can both be negative or overflow their respective limits, they will
// be normalized to a time between 00:00 - 23:59.
func New(hours, minutes int) Clock {
	nHours, nMins := normalize(hours, minutes, mph)
	_, nHours = normalize(day, nHours, hpd)
	return Clock{
		hour:   nHours,
		minute: nMins,
	}
}

// String returns the hours and minutes of a clock in "02:11" format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to a Clock. If the minutes overflow the remainder will be
// given to the hours. If the hours overflow the clock will start back at 0.
func (c Clock) Add(minutes int) Clock {
	c.hour, c.minute = normalize(c.hour, c.minute+minutes, mph)
	_, c.hour = normalize(day, c.hour, hpd)
	return c
}

// Subtract subtracts minutes from a Clock. If the minutes overflow the
// remainder will be given to the hours. If the hours overflow the clock will
// start back at 0.
func (c Clock) Subtract(minutes int) Clock {
	c.hour, c.minute = normalize(c.hour, c.minute-minutes, mph)
	_, c.hour = normalize(day, c.hour, hpd)
	return c
}

// normalize returns nhi, nlo such that
// hi * base + lo == nhi * base + nlo
// and nlo is between 0 and the base provided.
func normalize(hi, lo int, base timeRate) (nhi, nlo int) {
	rate := int(base)
	if lo < 0 {
		n := (-lo-1)/rate + 1
		hi -= n
		lo += n * rate
	}
	if lo >= rate {
		n := lo / rate
		hi += n
		lo -= n * rate
	}
	return hi, lo
}
