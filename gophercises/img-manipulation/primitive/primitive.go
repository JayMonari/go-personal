package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Mode int

const (
	combo Mode = iota
	triangle
	rect
	elipse
)

// WithMode is an option for the Transform function that will define the mode
// you want to use. By default
func WithMode(n int) func() []string {
	return func() []string {
		return []string{"-n", fmt.Sprintf("%d", n)}
	}
}

func Transform(image io.Reader, numShapes int, opts ...func() []string) (io.Reader, error) {
	// TODO(jaymonari): The extension needs to be .png or .jpeg
	in, err := ioutil.TempFile("", "in_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	out, err := ioutil.TempFile("", "out_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, err
	}
	stdCombo, err := primitive(in.Name(), out.Name(), numShapes, rect)
	if err != nil {
		return nil, err
	}
	fmt.Println(stdCombo)

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argStr)...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
