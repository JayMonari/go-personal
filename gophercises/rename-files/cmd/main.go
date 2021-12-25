package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}

func main() {
	dir := "sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _, origfn := range toRename {
		var newf file
		var err error
		newf.name, err = match(origfn.name)
		origPath := filepath.Join(dir, origfn.name)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newf.name)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
	}
}

func match(fn string) (string, error) {
	extIdx := strings.LastIndex(fn, ".")
	ext := fn[extIdx+1:]
	idIdx := strings.LastIndex(fn[:extIdx], "_")
	number, err := strconv.Atoi(fn[idIdx+1 : extIdx])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(fn[:idIdx]), number, ext), nil
}
