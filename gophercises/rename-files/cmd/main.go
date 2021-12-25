package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^(.+?) ([0-9]{4}) \(([0-9]+)\)\.(.+?)$`)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "Whether or not to perform a dry run.")
	flag.Parse()

	toRename := make(map[string][]string)
	filepath.Walk("sample", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		currDir := filepath.Dir(path)
		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(currDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})

	for key, files := range toRename {
		n := len(files)
		dir := filepath.Dir(key)
		sort.Strings(files)
		for i, fname := range files {
			res, _ := match(fname)
			newFname := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, fname)
			newPath := filepath.Join(dir, newFname)
			if dry {
				fmt.Printf("mv %s -> %s\n", oldPath, newPath)
			} else {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, newPath, err.Error())
				}
			}
		}
	}
}

type matchResult struct {
	base, ext string
	index     int
}

func match(fn string) (matchResult, error) {
	extIdx := strings.LastIndex(fn, ".")
	ext := fn[extIdx+1:]
	idIdx := strings.LastIndex(fn[:extIdx], "_")
	number, err := strconv.Atoi(fn[idIdx+1 : extIdx])
	if err != nil {
		return matchResult{}, err
	}
	return matchResult{
		base:  strings.Title(fn[:idIdx]),
		ext:   ext,
		index: number,
	}, nil
}
