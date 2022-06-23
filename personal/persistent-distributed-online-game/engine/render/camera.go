package render

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	// Position in world space where camera is.
	Pos pixel.Vec
	// Zoom is how much to zoom world space to screen space.
	Zoom float64
	// Camera operates on this to know what bounds are.
	win *pixelgl.Window
	// The full transformation the camera will apply to drawn things.
	mat pixel.Matrix
}

func NewCamera(win *pixelgl.Window, x, y float64) *Camera {
	return &Camera{
		win:  win,
		Pos:  pixel.V(x, y),
		Zoom: 1.0,
		mat:  pixel.IM,
	}
}

func (c *Camera) Update() {
	center := c.win.Bounds().Center()
	movePos := pixel.V(
		math.Floor(-c.Pos.X),
		math.Floor(-c.Pos.Y)).
		Add(center)
	c.mat = pixel.IM.Moved(movePos).Scaled(center, c.Zoom)
}

func (c *Camera) Mat() pixel.Matrix { return c.mat }
