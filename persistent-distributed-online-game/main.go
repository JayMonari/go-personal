package main

//go:generate packer --input images --stats

import (
	_ "image/png"
	"mmo/engine/asset"
	"mmo/engine/ecs"
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

	engine := ecs.NewEngine()

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
	spawnPoint := Transform{
		float64(tileSize * mapSize / 2),
		float64(tileSize * mapSize / 2),
	}
	manSprite, err := ss.Get("man1.png")
	check(err)
	hatmanSprite, err := ss.Get("man2.png")
	check(err)

	manID := engine.NewID()
	ecs.Write(engine, manID, manSprite)
	ecs.Write(engine, manID, spawnPoint)
	ecs.Write(engine, manID, Keybinds{
		Left:  pixelgl.KeyLeft,
		Right: pixelgl.KeyRight,
		Down:  pixelgl.KeyDown,
		Up:    pixelgl.KeyUp,
	})
	hatmanID := engine.NewID()
	ecs.Write(engine, hatmanID, hatmanSprite)
	ecs.Write(engine, hatmanID, spawnPoint)
	ecs.Write(engine, hatmanID, Keybinds{
		Left:  pixelgl.KeyA,
		Right: pixelgl.KeyD,
		Down:  pixelgl.KeyS,
		Up:    pixelgl.KeyW,
	})

	camera := render.NewCamera(win, 0, 0)
	zoomSpeed := 0.1
	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		if amt := win.MouseScroll().Y; amt != 0 {
			camera.Zoom += zoomSpeed * amt
		}

		HandleInput(win, engine)
		t := Transform{}
		if ok := ecs.Read(engine, manID, &t); ok {
			camera.Pos = pixel.V(t.X, t.Y)
		}
		camera.Update()

		win.SetMatrix(camera.Mat())
		// Collision Detection would go here.
		tmap.Draw(win)
		DrawSprites(win, engine)

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

func (t *Keybinds) ComponentSet(val any) { *t = val.(Keybinds) }

type Sprite struct{ *pixel.Sprite }

func (t *Sprite) ComponentSet(val any) { *t = val.(Sprite) }

type Transform struct{ X, Y float64 }

func (t *Transform) ComponentSet(val any) { *t = val.(Transform) }

func DrawSprites(win *pixelgl.Window, e *ecs.Engine) {
	ecs.Each(e, Sprite{}, func(id ecs.ID, a any) {
		s := a.(Sprite)
		t := Transform{}
		ok := ecs.Read(e, id, &t)
		if !ok {
			return
		}
		s.Draw(win, pixel.IM.Scaled(pixel.ZV, 1.0).Moved(pixel.V(t.X, t.Y)))
	})
}

func HandleInput(win *pixelgl.Window, e *ecs.Engine) {
	put := win.Pressed
	ecs.Each(e, Keybinds{}, func(id ecs.ID, a any) {
		kb := a.(Keybinds)
		t := Transform{}
		ok := ecs.Read(e, id, &t)
		if !ok {
			return
		}
		if put(kb.Left) {
			t.X -= 2.0
		}
		if put(kb.Right) {
			t.X += 2.0
		}
		if put(kb.Down) {
			t.Y -= 2.0
		}
		if put(kb.Up) {
			t.Y += 2.0
		}
		ecs.Write(e, id, t)
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
