package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fName := "birthday_001.txt"
	// newName, err := match(fName, 4)
	// if err != nil {
	// 	panic("no match")
	// }
	// fmt.Println(newName)
	dir := "sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
	}
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
