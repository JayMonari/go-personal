package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// This package has been deprecated but to understand what might be seen as
// utilities we'll flesh it out here and move it later.
func main() {
	ExampleWriteFile()
	ExampleReadDir()
}

// Wrapper for functions that require ReadClosers, but you only have a Reader.
func NopCloser() {
  io.NopCloser(strings.NewReader("some string")) // io.ReadCloser
}

// When is this useful? Testing maybe?
func ExampleReadAll() {
	r := strings.NewReader("NewReader returns a new Reader reading from s.\nIt is similar to bytes.NewBufferString but more efficient and read-only.")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

// Seems like a very useful utility for walking.
func ExampleReadDir() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println("FILE NAME:", f.Name())
	}
}

// Being able to stream a file so that you may change the file maybe?
func ExampleReadFile() {
	f, err := os.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(f))
}

// calls os.MkdirTemp
func ExampleTempDir() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir(".", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	fmt.Println(os.TempDir())

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

// Now calls os.CreateTemp
func ExampleTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile(".", "example")
	if err != nil {
		log.Fatal(err)
	}
  defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
}

// calls os.WriteFile
func ExampleWriteFile() {
  msg := []byte("Hello, Gophers!")
  if err := os.WriteFile("newfile", msg, 0666); err != nil {
    log.Fatal(err)
  }
}
