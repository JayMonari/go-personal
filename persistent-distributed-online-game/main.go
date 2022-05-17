package main

//go:generate packer --input images --stats

import (
	_ "image/png"
	"mmo/engine/asset"
	"mmo/engine/pgen"
	"mmo/engine/render"
	"mmo/engine/tilemap"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(runGame)
}

const (
	GrassTile tilemap.TileType = iota
	DirtTile
	WaterTile
)

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title:     "MMO",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	check(err)
	win.SetSmooth(false)

	ss, err := asset.NewLoad(os.DirFS("./")).SpriteSheet("packed.json")
	check(err)

	// Create Tilemap
	waterLevel := 0.5
	beachLevel := waterLevel + 0.1
	terrain := pgen.NewNoiseMap(time.Now().UnixNano(), 1.0)
	tileSize := 16
	mapSize := 500
	tiles := make([][]tilemap.Tile, mapSize, mapSize)
	for x := range tiles {
		tiles[x] = make([]tilemap.Tile, mapSize, mapSize)
		for y := range tiles[x] {
			if height := terrain.Get(x, y); height < waterLevel {
				tiles[x][y] = GetTile(ss, WaterTile)
			} else if height < beachLevel {
				tiles[x][y] = GetTile(ss, DirtTile)
			} else {
				tiles[x][y] = GetTile(ss, GrassTile)
			}
		}
	}
	tmap := tilemap.New(
		tiles,
		tileSize,
		pixel.NewBatch(&pixel.TrianglesData{}, ss.Picture()),
	)
	tmap.Rebatch()

	// Creaet people
	spawnPoint := pixel.V(
		float64(tileSize*mapSize/2),
		float64(tileSize*mapSize/2),
	)
	manSprite, err := ss.Get("man1.png")
	check(err)
	hatmanSprite, err := ss.Get("man2.png")
	check(err)
	people := make([]*Person, 0, 2)
	people = append(people,
		NewPerson(manSprite, spawnPoint, Keybinds{
			Left:  pixelgl.KeyLeft,
			Right: pixelgl.KeyRight,
			Down:  pixelgl.KeyDown,
			Up:    pixelgl.KeyUp,
		}),
		NewPerson(hatmanSprite, spawnPoint, Keybinds{
			Left:  pixelgl.KeyA,
			Right: pixelgl.KeyD,
			Down:  pixelgl.KeyS,
			Up:    pixelgl.KeyW,
		}),
	)

	camera := render.NewCamera(win, 0, 0)
	zoomSpeed := 0.1
	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		if amt := win.MouseScroll().Y; amt != 0 {
			camera.Zoom += zoomSpeed * amt
		}

		for _, p := range people {
			p.HandleInput(win)
		}
		camera.Pos = people[0].Position
		camera.Update()

		win.SetMatrix(camera.Mat())
		// Collision Detection would go here.
		tmap.Draw(win)
		for _, p := range people {
			p.Draw(win)
		}
		win.SetMatrix(pixel.IM)

		win.Update()
	}
}

func GetTile(ss *asset.SpriteSheet, t tilemap.TileType) tilemap.Tile {
	name := ""
	switch t {
	case GrassTile:
		name = "grass.png"
	case DirtTile:
		name = "dirt.png"
	case WaterTile:
		name = "water.png"
	default:
		panic("Unknown TileType!")
	}
	s, err := ss.Get(name)
	check(err)
	return tilemap.Tile{Type: t, Sprite: s}
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
	put := win.Pressed
	if put(p.Keybinds.Left) {
		p.Position.X -= 2.0
	}
	if put(p.Keybinds.Right) {
		p.Position.X += 2.0
	}
	if put(p.Keybinds.Down) {
		p.Position.Y -= 2.0
	}
	if put(p.Keybinds.Up) {
		p.Position.Y += 2.0
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
