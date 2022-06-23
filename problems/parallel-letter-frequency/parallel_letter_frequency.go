package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune of each string and
// returns the rune count in a FreqMap.
func ConcurrentFrequency(strings []string) FreqMap {
	freqs := make(chan FreqMap)
	defer close(freqs)

	for _, s := range strings {
		go func(s string) {
			freqs <- Frequency(s)
		}(s)
	}
	conFM := <-freqs
	for i := 1; i < len(strings); i++ {
		for rn, cnt := range <-freqs {
			conFM[rn] += cnt
		}
	}
	return conFM
}
