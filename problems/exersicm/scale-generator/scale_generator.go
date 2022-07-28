package scale

import "strings"

type chromaticScale [12]string

const (
	chrLen = 12
)

var (
	chrSharp   = chromaticScale{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	chrFlat    = chromaticScale{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	flatTonics = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	intervals  = map[rune]int{
		'm': 1,
		'M': 2,
		'A': 3,
	}
)

// Scale returns a musical scale starting with the tonic and following the
// specified interval pattern. If interval is empty a full chromatic (12-note)
// scale will be returned.
func Scale(tonic, interval string) []string {
	chrScale := pickChrScale(tonic)
	if interval == "" {
		return buildChromaticScale(tonic, chrScale)
	}

	scale := make([]string, len(interval))
	note := findStart(tonic, chrScale)
	for i, step := range interval {
		scale[i] = chrScale[note]
		note = (note + intervals[step]) % chrLen
	}
	return scale
}

// pickChrScale returns the correct chromatic chrScale based on the given
// tonic.
func pickChrScale(tonic string) chromaticScale {
	for _, note := range flatTonics {
		if note == tonic {
			return chrFlat
		}
	}
	return chrSharp
}

// buildChromaticScale returns the chromatic scale shifted based on the tonic.
func buildChromaticScale(tonic string, chrScale chromaticScale) []string {
	note := findStart(tonic, chrScale)
	scale := chrScale[note:]
	return append(scale, chrScale[:note]...)
}

// findStart returns the index of the tonic in the scale. If the tonic is
// not in the chrScale -1 is returned.
func findStart(tonic string, chrScale chromaticScale) int {
	tonic = strings.ToUpper(tonic[:1]) + tonic[1:]
	for i, note := range chrScale {
		if note == tonic {
			return i
		}
	}
	return -1
}
