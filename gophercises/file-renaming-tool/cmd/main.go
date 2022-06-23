package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile(`^(.+?) (\d{4}) \((\d+) of (\d+)\)\.(.+?)$`)

const replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "Whether or not to perform a dry run.")
	flag.Parse()

	for _, oldPath := range findFiles() {
		newFilename := re.ReplaceAllString(filepath.Base(oldPath), replaceString)
		newPath := filepath.Join(filepath.Dir(oldPath), newFilename)
		switch dry {
		case true:
			fmt.Printf("mv %s -> %s\n", oldPath, newPath)
		case false:
			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}
}

func findFiles() (toRename []string) {
	filepath.Walk("sample", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if re.MatchString(info.Name()) {
			toRename = append(toRename, path)
		}
		return nil
	})
	return toRename
}
