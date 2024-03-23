package primitive

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Mode defines the shapes used when transforming images.
type Mode interface {
	isMode()
}

type mode uint8

func (mode) isMode() {}

const (
	ModeCombo mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

func Transform(
	ctx context.Context,
	img interface {
		io.Reader
		Name() string
	},
	nShapes int,
	opts ...func() []string,
) (io.Reader, error) {
	out, err := os.CreateTemp("", "*"+filepath.Base(img.Name()))
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	fmt.Println(img.Name(), out.Name())
	raw, err := primitive(ctx, img.Name(), out.Name(), nShapes, ModeCombo)
	fmt.Println(string(raw))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.ReadFrom(out)
	return &buf, nil
}

func primitive(
	ctx context.Context, input, output string, nShapes int, m Mode,
) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "primitive", strings.Fields(
		fmt.Sprintf("-i %s -o %s -n %d -m %d", input, output, nShapes, m),
	)...)
	return cmd.CombinedOutput()
}
