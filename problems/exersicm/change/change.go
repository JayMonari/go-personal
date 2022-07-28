package change

import "errors"

// Change determines the fewest number of coins to be given to a customer such
// that the sum of the coins' value would equal the target.
func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	amounts := make([][]int, target+1)
	amounts[0] = []int{}
	for n := range amounts {
		for _, c := range coins {
			if n-c < 0 || amounts[n-c] == nil {
				continue
			}
			if amounts[n] == nil || len(amounts[n-c])+1 < len(amounts[n]) {
				amounts[n] = append([]int{c}, amounts[n-c]...)
			}
		}
	}
	if amounts[target] == nil {
		return nil, errors.New("no change could be made")
	}
	return amounts[target], nil
}
