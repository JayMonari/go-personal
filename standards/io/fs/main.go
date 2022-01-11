package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
  ValidPathExample()
}

func GlobExample() {
	myFS := os.DirFS("nested")
	matches1, err := fs.Glob(myFS, "file*")
	matches2, err := fs.Glob(myFS, "*go")
	if err != nil {
		panic(err)
	}
	for _, m := range matches1 {
		fmt.Println("A MATCH!!!!", m)
	}
	for _, m := range matches2 {
		fmt.Println("A MATCH!!!!", m)
	}
}

func ValidPathExample() {
  if fs.ValidPath("nested/file.txt") {
    fmt.Println("IT WAS VALID!!!")
  }
}
