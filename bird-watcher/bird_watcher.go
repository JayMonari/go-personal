package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, cnt := range birdsPerDay {
		sum += cnt
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	i := (week - 1) * 7
	sum := 0
	for _, cnt := range birdsPerDay[i : i+7] {
		sum += cnt
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, c := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = c + 1
		}
	}
	return birdsPerDay
}
