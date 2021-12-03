package main

import "fmt"

// This was dumb
func main() {
  var input string
  fmt.Scanf("%s\n", &input)
  answer := 1
  for _, r := range input {
    if r >= 'A' && r <= 'Z' {
      answer++
    }
  }
}
