package main

//go:generate packer --input images --stats

import (
	_ "image/png"
	"mmo/engine/asset"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title:     "MMO",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)

	load := asset.NewLoad(os.DirFS("./"))
	ss, err := load.SpriteSheet("packed.json")
	if err != nil {
		panic(err)
	}
	manSprite, err := ss.Get("man.png")
	if err != nil {
		panic(err)
	}
	manPos := win.Bounds().Center()
	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		if win.Pressed(pixelgl.KeyLeft) {
			manPos.X -= 2.0
		}
		if win.Pressed(pixelgl.KeyRight) {
			manPos.X += 2.0
		}
		if win.Pressed(pixelgl.KeyDown) {
			manPos.Y -= 2.0
		}
		if win.Pressed(pixelgl.KeyUp) {
			manPos.Y += 2.0
		}
		manSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(manPos))

		win.Update()
	}
}
