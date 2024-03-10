package primitive

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
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

func primitive(ctx context.Context, input, output string, nShapes int, m Mode) {
	cmd := exec.CommandContext(ctx, "primitive", strings.Fields(
		fmt.Sprintf("-i %s -o %s -n %d -m %d", input, output, nShapes, m),
	)...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		slog.Error("cmd.Run", "err", err)
		return
	}
}
