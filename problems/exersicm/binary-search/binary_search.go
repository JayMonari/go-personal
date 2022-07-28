package binarysearch

// SearchInts finds the index of a value t in a sorted slice of ints.
func SearchInts(nums []int, t int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi - lo)/2
		cand := nums[mid]
		switch {
		case cand < t:
			lo = mid + 1
		case cand > t:
			hi = mid - 1
		case cand == t:
			return mid
		}
	}
	return -1
}
