package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/go
func InArray(substrs []string, strs []string) (out []string) {
	subs := make(map[string]struct{})
	for _, ss := range substrs {
		for _, s := range strs {
			if strings.Contains(s, ss) {
				subs[ss] = struct{}{}
			}
		}
	}
	for s := range subs {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func main() {
	fmt.Println(InArray([]string{"live", "arp", "strong"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"}))
	// []string{"arp","live","strong"}

	fmt.Println(InArray([]string{"cod", "code", "wars", "ewar", "ar"},
		[]string{}))
	// []string{}
}
