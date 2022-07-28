package summultiples

// SumMultiples returns the sum of all the factors of each number in nums up to
// the provided limit.
func SumMultiples(limit int, nums ...int) int {
	sum := 0
	for m := range getMultiples(limit, nums) {
		sum += m
	}
	return sum
}

// getMultiples returns a set of all factors up to the limit for each number in
// nums.
func getMultiples(limit int, nums []int) map[int]struct{} {
	muls := map[int]struct{}{}
	for _, num := range nums {
		if num == 0 {
			continue
		}
		for n := num; n < limit; n += num {
			muls[n] = struct{}{}
		}
	}
	return muls
}
