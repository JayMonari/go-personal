package slice

import "fmt"

// Basic shows how to create a slice and how to set and get values in it.
func Basic() {
	// Don't add a number between the
	// `[]` brackets and we `make` slices if we
	// want to have a capacity and length
	slice := make([]string, 3)
	fmt.Println("empty:", slice)

	slice[0] = "|set zeroeth value|"
	slice[1] = "|set first value|"
	slice[2] = "|set second value|"
	fmt.Println("full:", slice)
	fmt.Println("pick a value:", slice[2])
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))

	inline := []int{0, 1, 2, 3, 4}
	fmt.Println("Can be declared inline", inline)
}

// Append shows how to put more elements into a slice even if we don't
// have the capacity for it using `append`.
func Append() {
	// Why wouldn't I do this always?
	var slice []string
	// Good Question! Lets answer it!
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))

	slice = append(slice, "append a single value")
	slice = append(slice, "append", "multiple", "values")
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))
	fmt.Println("We had to go find more space! Which takes time and effort!")
	fmt.Println("slice:", slice)

	unpackAllThese := []string{"`...`", "is used to put", "all the values in", "at the same time"}
	slice = append(slice, unpackAllThese...)
	fmt.Println("capacity:", cap(slice))
	fmt.Println("length:", len(slice))
	fmt.Println("We had to go find even more space!!!")
	fmt.Println("slice:", slice)
}

// Copy shows how to copy one slice into another slice using the builtin
// `copy` function.
func Copy() {
	// src is short for source
	srcSlice := make([]int, 10)
	fmt.Println("empty srcSlice:", srcSlice)
	for i := 0; i < 10; i++ {
		srcSlice[i] = i
	}
	fmt.Println("full srcSlice:", srcSlice)

	// dst is short for destination
	dstSlice := make([]int, len(srcSlice))
	fmt.Println("empty dstSlice:", dstSlice)
	copy(dstSlice, srcSlice)
	fmt.Println("full dstSlice:", dstSlice)
}

// IndexOutOfRangePanic shows us what happens when we try to access an
// index that does not exist in a slice.
func IndexOutOfRangePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("slice paniced!\n", r)
		}
	}()

	sl := make([]int, 5)
	// Change -1 to 0 to see the panic happen at the other end of the slice.
	for i := -1; i < len(sl)+1; i++ {
		fmt.Println("NOTE(jay): this is going to panic before we ever see this!", sl[i])
	}
}

// ReasonForName shows us why a slice is called a slice and that's because we
// can take slices (pieces) of a slice depending on our needs using the `:`
// slice operator.
func ReasonForName() {
	var slice = []string{"zero", "one", "two", "three", "four", "five"}
	fmt.Printf("sliceUpToThirdIndex: %v\nlength: %d capacity: %d\n",
		slice,
		len(slice),
		cap(slice))

	sliceUpToThirdIndex := slice[:3]
	fmt.Printf("sliceUpToThirdIndex: %v\nlength: %d capacity: %d\n",
		sliceUpToThirdIndex,
		len(sliceUpToThirdIndex),
		cap(sliceUpToThirdIndex))

	sliceStartAtIndexTwo := slice[2:]
	fmt.Printf("sliceStartAtIndexTwo: %v\nlength: %d capacity: %d\n",
		sliceStartAtIndexTwo,
		len(sliceStartAtIndexTwo),
		cap(sliceStartAtIndexTwo))

	sliceFromOneUpToFour := slice[1:4]
	fmt.Printf("sliceFromOneUpToFour: %v\nlength: %d capacity: %d\n",
		sliceFromOneUpToFour,
		len(sliceFromOneUpToFour),
		cap(sliceFromOneUpToFour))
	s := "Max Efficiency"
	fmt.Println(s[4:], "to the", s[:3], "for substrings")
}

// Matrix shows how to make a matrix also known as a 2d array, but still
// have the flexibility of slices!
func Matrix() {
	// We will allocate three slices in a slice
	matrix := make([][]int, 3)
	fmt.Println("matrix empty:", matrix)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matrix[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix full:", matrix)

	// and we can treat each slice like we would any other slice.
	matrix[0] = append(matrix[0], 21)
	fmt.Println("matrix append first slice with value:", matrix)
}
