package twelve

import (
	"fmt"
	"strings"
)

var verses = [amtDays]string{
	"twelve Drummers Drumming,",
	"eleven Pipers Piping,",
	"ten Lords-a-Leaping,",
	"nine Ladies Dancing,",
	"eight Maids-a-Milking,",
	"seven Swans-a-Swimming,",
	"six Geese-a-Laying,",
	"five Gold Rings,",
	"four Calling Birds,",
	"three French Hens,",
	"two Turtle Doves,",
	"and a Partridge in a Pear Tree.",
}

var days = [amtDays]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

const (
	verseFmt = "On the %s day of Christmas my true love gave to me: %s"
	amtDays  = 12
)

// Song returns entire song with each verse ending with a newline.
func Song() string {
	verses := make([]string, amtDays)
	for i := 1; i <= amtDays; i++ {
		verses[i-1] = Verse(i)
	}
	return strings.Join(verses, "\n")
}

// Verse returns verse n of song.
func Verse(n int) string {
	if n == 1 {
		return fmt.Sprintf(verseFmt, days[n-1], verses[amtDays-1][4:])
	}
	return fmt.Sprintf(verseFmt, days[n-1], strings.Join(verses[amtDays-n:], " "))
}
