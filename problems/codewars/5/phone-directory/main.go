package main

import (
	"regexp"
	"strings"
)

// https://www.codewars.com/kata/56baeae7022c16dd7400086e/train/go
func Phone(dir, num string) string {
	cnt := 0
	data := ""
	for _, dd := range strings.Split(dir, "\n") {
		if strings.Contains(dd, "+"+num) {
			cnt++
			data = dd
		}
	}
	if cnt == 0 {
		return "Error => Not found: " + num
	} else {
		if cnt > 1 {
			return "Error => Too many people: " + num
		}
	}
	reg, _ := regexp.Compile("<|>")
	name := reg.Split(data, -1)[1]
	reg, _ = regexp.Compile(num + "|<.*>|[^A-Za-z0-9\\. -]")
	address := string(reg.ReplaceAll([]byte(data), []byte(" ")))
	reg, _ = regexp.Compile("\\s+")
	address = string(reg.ReplaceAll([]byte(address), []byte(" ")))
	address = strings.Trim(address, " ")
	return "Phone => " + num + ", Name => " + name + ", Address => " + address
}
