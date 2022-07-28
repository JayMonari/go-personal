package graphs_test

import "graphs"


var graph = map[string][]string{
	"a": {"c", "b"},
	"b": {"d"},
	"c": {"e"},
	"d": {"f"},
	"e": {},
	"f": {},
}

func ExampleDepthFirstPrint() {
	graphs.DepthFirstPrint(graph, "a")
	// Output:
	// a
	// b
	// d
	// f
	// c
	// e
}

func ExampleBreadthFirstPrint() {
	graphs.BreadthFirstPrint(graph, "a")
	// Output:
	// a
	// c
	// b
	// e
	// d
	// f
}
