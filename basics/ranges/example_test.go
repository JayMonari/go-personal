package ranges_test

import "basics/ranges"

func ExampleRangeIndex() {
	ranges.RangeIndex()
	// Output:
	// index: 0
	// index: 1
	// index: 2
	// index: 3
	// index: 4
	// index: 5
	// index: 6
	// index: 7
	// index: 8
	// index: 9
}

func ExampleRangeIndexAndValues() {
	ranges.RangeIndexAndValues()
	// Output:
	// index: 0, access value: 1, range value: 1
	// index: 1, access value: 2, range value: 2
	// index: 2, access value: 3, range value: 3
	// index: 3, access value: 4, range value: 4
	// index: 4, access value: 5, range value: 5
	// nums: [1 4 9 16 25]
}

func ExampleRangeValues() {
	ranges.RangeValues()
	// Output:
	// friend: Gabby
	// friend: Gorm
	// friend: Gunter
}

func ExampleRangeMap() {
	// XXX: This may fail from time to time!
	// There is no order in maps!
	// Why not run it a few times to see? 🙂
	ranges.RangeMap()
	// Output:
	// Gaph is married.
	// Gene is not married.
	// Gable is not married.
}

func ExampleRangeString() {
	ranges.RangeString()
	// Output:
	// index: 0 rune: g
	// index: 1 rune: o
	// index: 2 rune: p
	// index: 3 rune: h
	// index: 4 rune: e
	// index: 5 rune: r
	// index: 6 rune: g
	// index: 7 rune: o
	// index: 8 rune: .
	// index: 9 rune: d
	// index: 10 rune: e
	// index: 11 rune: v
}
