package gigasecond

import "time"

// AddGigasecond returns a Time 1,000,000,000 seconds in the future
func AddGigasecond(t time.Time) time.Time {
	d, _ := time.ParseDuration("1000000000s")
	return t.Add(d)
}
