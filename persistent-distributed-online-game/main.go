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
	manSprite, err := ss.Get("man1.png")
	if err != nil {
		panic(err)
	}
	manPos := win.Bounds().Center()
	hatmanSprite, err := ss.Get("man2.png")
	if err != nil {
		panic(err)
	}
	hatmanPos := win.Bounds().Center()
	people := make([]*Person, 0, 2)
	people = append(people,
		NewPerson(manSprite, manPos, Keybinds{
			Left:  pixelgl.KeyLeft,
			Right: pixelgl.KeyRight,
			Down:  pixelgl.KeyDown,
			Up:    pixelgl.KeyUp,
		}),
		NewPerson(hatmanSprite, hatmanPos, Keybinds{
			Left:  pixelgl.KeyA,
			Right: pixelgl.KeyD,
			Down:  pixelgl.KeyS,
			Up:    pixelgl.KeyW,
		}),
	)
	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		for _, p := range people {
			p.HandleInput(win)
		}
		// Collision Detection would go here.
		for _, p := range people {
			p.Draw(win)
		}

		win.Update()
	}
}

type Keybinds struct {
	Up, Down, Left, Right pixelgl.Button
}

type Person struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
	Keybinds
}

func NewPerson(s *pixel.Sprite, pos pixel.Vec, k Keybinds) *Person {
	return &Person{
		Sprite:   s,
		Position: pos,
		Keybinds: k,
	}
}

func (p *Person) Draw(win *pixelgl.Window) {
	p.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(p.Position))
}

func (p *Person) HandleInput(win *pixelgl.Window) {
	if win.Pressed(p.Keybinds.Left) {
		p.Position.X -= 2.0
	}
	if win.Pressed(p.Keybinds.Right) {
		p.Position.X += 2.0
	}
	if win.Pressed(p.Keybinds.Down) {
		p.Position.Y -= 2.0
	}
	if win.Pressed(p.Keybinds.Up) {
		p.Position.Y += 2.0
	}
}
