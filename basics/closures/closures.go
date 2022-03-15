package closure

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
)

// ClosureAsGenerator
func ClosureAsGenerator() func() int {
	startNum := 0
	return func() int {
		startNum++
		return startNum
	}
}

// ChangeValues
type ChangeValues struct {
	MyStr     string
	IsChanged bool
}

// ClosureAsMiddleware
func ClosureAsMiddleware(myFunc func(cv *ChangeValues, n *int)) func(cv *ChangeValues, n *int) {
	return func(cv *ChangeValues, n *int) {
		fmt.Printf("this is a statement that happens **BEFORE** myFunc:\n%+v and %d", cv, n)
		myFunc(cv, n)
		fmt.Printf("this is a statement that happens **AFTER** myFunc:\n%+v and %d", cv, n)
	}
}

// ClosureForAccessingData
func ClosureForAccessingData(cv ChangeValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, cv.MyStr)
	}
}

// ClosureForStandardLibrary
func ClosureForStandardLibrary() {
	fullOfXs := "XSXXXEXXCRXEXXXXTXXHXEXRXEXX"
	mappedStr := strings.Map(func(r rune) rune {
		if r == 'X' {
			return ' '
		}
		return r
	}, fullOfXs)
	fmt.Println(mappedStr)

	sortedInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	index := sort.Search(len(sortedInts), func(i int) bool {
		return sortedInts[i] < 14
	})
	fmt.Println("index found from Binary Search:", index)
}

var doWork1 = func(args ...int) int {
	sum := 0
	for _, n := range args {
		sum += n
	}
	return sum
}

var doWork2 = func(sum int) float32 {
	crunchNums := math.Pi * 2
	return float32(sum) * float32(crunchNums)
}

var doWork3 = func(floaty float32) string {
	return fmt.Sprintf("%0.4f", floaty)
}

// ClosureAvoidCallbackHell
func ClosureAvoidCallbackHell() {
	result1 := doWork1(1, 2, 3, 4, 5, 6, 7, 8)
	result2 := doWork2(result1)
	result3 := doWork3(result2)
	fmt.Printf("Some padding for final result:\n%-10s", result3)
}
