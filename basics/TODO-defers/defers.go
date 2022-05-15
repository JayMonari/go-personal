package defers

import (
	"fmt"
	"os"
)

// defer runs after return
// defer for closing file
// defer os.Remove/All(tmp.Name()/tDir)
// defer for closing HTML body
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// defer cancel call with context
// defer for recovery
