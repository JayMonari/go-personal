package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fName := "birthday_001.txt"
	newName, err := match(fName, 4)
	if err != nil {
		panic("no match")
	}
	fmt.Println(newName)
}

func match(fn string, total int) (string, error) {
	extIdx := strings.LastIndex(fn, ".")
	ext := fn[extIdx+1:]
	idIdx := strings.LastIndex(fn[:extIdx], "_")
	number, err := strconv.Atoi(fn[idIdx+1 : extIdx])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(fn[:idIdx]), number, total, ext), nil
}
