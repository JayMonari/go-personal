package array

import "fmt"

// Arrays shows all the ways to make and manipulate arrays in Go.
func Arrays() {
	var arr [4]int
	fmt.Println("Length of arr:", len(arr))
	fmt.Println("Empty:", arr)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	fmt.Println("Filled:", arr)
	fmt.Println("Grab one value:", arr[3])

	inline := [4]int{2, 3, 4, 5}
	fmt.Println("Can be declared inline", inline)

	// If you don't want to count how many values, but you don't want a slice.
	// you can use the `...` syntax which will make an array of however many
	// values that you initialize it with.
	constSlice := [...]int{8, 9, 10, 11, 15}
	fmt.Printf("Just an array %T\nValues:%v", constSlice, constSlice)
}

// Matrix shows how to make a 2-dimensional array aka a matrix in Go.
func Matrix() {
	var matrix [4][5]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("Matrix:", matrix)
}
