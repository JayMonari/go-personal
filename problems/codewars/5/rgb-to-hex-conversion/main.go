package main

import "fmt"

// https://www.codewars.com/kata/513e08acc600c94f01000001/train/go
func RGB(r, g, b int) string {
	validate(&r, &g, &b)
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func validate(v ...*int) {
	for _, n := range v {
		if *n < 0 {
			*n = 0
		} else if *n > 255 {
			*n = 255
		}
	}
}

func main() {
	fmt.Println(
		RGB(255, -255, 10000),
		RGB(0, 255, 255),
		RGB(0, 0, 255),
		RGB(0, 0, 0),
		RGB(255, 0, 0),
		RGB(140, 0, 211),
	)
}
