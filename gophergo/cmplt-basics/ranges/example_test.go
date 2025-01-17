package ranges_test

import "basics/ranges"

func ExampleIndex() {
	ranges.Index()
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

func ExampleValues() {
	ranges.Values()
	// Output:
	// friend: Gabby
	// friend: Gorm
	// friend: Gunter
}

func ExampleIndexAndValues() {
	ranges.IndexAndValues()
	// Output:
	// index: 0, access value: 1, range value: 1
	// index: 1, access value: 2, range value: 2
	// index: 2, access value: 3, range value: 3
	// index: 3, access value: 4, range value: 4
	// index: 4, access value: 5, range value: 5
	// nums: [1 4 9 16 25]
}

func ExampleMap() {
	// XXX(jay): This may fail from time to time!
	// There is no order in maps!
	// Why not run it a few times to see? 🙂
	ranges.Map()
	// Output:
	// Gaph is married.
	// Gene is not married.
	// Gable is not married.
}

func ExampleString() {
	ranges.String()
	// Output:
	// index: 0 rune: 103 representation: g
	// index: 1 rune: 111 representation: o
	// index: 2 rune: 112 representation: p
	// index: 3 rune: 104 representation: h
	// index: 4 rune: 101 representation: e
	// index: 5 rune: 114 representation: r
	// index: 6 rune: 103 representation: g
	// index: 7 rune: 111 representation: o
	// index: 8 rune: 46 representation: .
	// index: 9 rune: 100 representation: d
	// index: 10 rune: 101 representation: e
	// index: 11 rune: 118 representation: v
}

func ExampleChannel() {
	ranges.Channel()
	// Output:
	// We can get
	// values from a channel
	// continuously.
	// Just make sure
	// you close the channel
	// at some time 😉
}

func ExampleScopedValues() {
	ranges.ScopedValues()
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
