package main

import "fmt"

// https://www.codewars.com/kata/514b92a657cdc65150000006/train/go
func Multiple3And5(limit int) (sum int) {
	seen := map[int]struct{}{}
	for i := 3; i < limit; i++ {
		if _, ok := seen[i]; ok {
			continue
		}
		if i%3 != 0 && i%5 != 0 {
			continue
		}
		seen[i] = struct{}{}
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(Multiple3And5(10))
}
