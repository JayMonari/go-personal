package primitive

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Mode defines the shapes used when transforming images.
type Mode interface {
	fmt.Stringer
	Val() int

	isMode()
}

type mode uint8

func (mode) isMode() {}

func (m mode) Val() int { return int(m) }

func (m mode) String() string {
	switch m {
	case ModeCombo:
		return "combo"
	case ModeTriangle:
		return "triangle"
	case ModeRect:
		return "rect"
	case ModeEllipse:
		return "ellipse"
	case ModeCircle:
		return "circle"
	case ModeRotatedRect:
		return "rotatedRect"
	case ModeBeziers:
		return "beziers"
	case ModeRotatedEllipse:
		return "rotatedEllipse"
	case ModePolygon:
		return "polygon"
	default:
		return "unknown"
	}
}

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

type ImageFD struct {
	io.Reader
	Ext string
}

func Transform(
	ctx context.Context,
	img ImageFD,
	nShapes int,
	mode Mode,
) (io.Reader, error) {
	in, err := os.CreateTemp("", "*"+img.Ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	if _, err := io.Copy(in, img); err != nil {
		return nil, err
	}

	out, err := os.CreateTemp("", "*"+img.Ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	raw, err := primitive(ctx, in.Name(), out.Name(), nShapes, mode)
	if len(raw) > 0 {
		fmt.Println(string(raw))
	}
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
