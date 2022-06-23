package graphs

import "fmt"

func DepthFirstPrint(graph map[string][]string, source string) {
	stk := []string{source}
	for len(stk) > 0 {
		vert := stk[len(stk)-1]
		fmt.Println(vert)
		stk = stk[:len(stk)-1]
		stk = append(stk, graph[vert]...)
	}
}

func BreadthFirstPrint(graph map[string][]string, source string) {
	q := []string{source}
	for len(q) > 0 {
		vert := q[0]
		q = q[1:]
		fmt.Println(vert)
		q = append(q, graph[vert]...)
	}
}
