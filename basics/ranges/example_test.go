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

func ExampleRangeChannel() {
	ranges.RangeChannel()
	// Output:
	// We can get
	// values from a channel
	// continuously.
	// Just make sure
	// you close the channel
	// at some time 😉
}

func ExampleRangeScopedValues() {
	ranges.RangeScopedValues()
	// Output:
	// Try to change by just the value
	// before: 0 after: 9
	// Never changes: [0 1 2 3 4]
	// before: 1 after: 9
	// Never changes: [0 1 2 3 4]
	// before: 2 after: 9
	// Never changes: [0 1 2 3 4]
	// before: 3 after: 9
	// Never changes: [0 1 2 3 4]
	// before: 4 after: 9
	// Never changes: [0 1 2 3 4]
	// Same! [0 1 2 3 4]
	// before: 78 after: 88
	// before: 79 after: 88
	// before: 84 after: 88
	// before: 32 after: 88
	// before: 99 after: 88
	// before: 104 after: 88
	// before: 97 after: 88
	// before: 110 after: 88
	// before: 103 after: 88
	// before: 101 after: 88
	// before: 100 after: 88
	// Same! NOT changed
	// before: true after: false
	// before: true after: false
	// before: true after: false
	// Same! map[x:true y:true z:true]
	//
	// Change by index (by dereference)
	// Changed! [9 9 9 9 4]
	// Changed! NOT XXXXXXX
	// Changed! map[x:false y:false z:false]
}
