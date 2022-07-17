package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running = true
	bkgClr  = rl.NewColor(147, 211, 196, 255)

	grassSprite  rl.Texture2D
	hillSprite   rl.Texture2D
	fenceSprite  rl.Texture2D
	houseSprite  rl.Texture2D
	waterSprite  rl.Texture2D
	tilledSprite rl.Texture2D

	plrSprite                         rl.Texture2D
	plrSrc                            rl.Rectangle
	plrDest                           rl.Rectangle
	plrSpeed                          float32 = 3
	plrMoving                         bool
	plrDir                            int
	plrUp, plrDown, plrRight, plrLeft bool
	plrFrame                          int

	frameCount int

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D

	mapW    int
	mapH    int
	tileMap []int
	srcMap  []byte

	tex      rl.Texture2D
	tileSrc  rl.Rectangle
	tileDest rl.Rectangle
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sprouts Lands")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/tilesets/Grass.png")
	hillSprite = rl.LoadTexture("res/tilesets/Hills.png")
	fenceSprite = rl.LoadTexture("res/tilesets/Building_parts/Fences.png")
	houseSprite = rl.LoadTexture("res/tilesets/Building_parts/Wooden_House.png")
	waterSprite = rl.LoadTexture("res/tilesets/Water.png")
	tilledSprite = rl.LoadTexture("res/tilesets/Tilled_Dirt.png")

	plrSprite = rl.LoadTexture("res/characters/Basic_Charakter_Spritesheet.png")

	plrSrc = rl.NewRectangle(0, 0, 48, 48)
	plrDest = rl.NewRectangle(200, 200, 100, 100)

	tileSrc = rl.NewRectangle(0, 0, 16, 16)
	tileDest = rl.NewRectangle(0, 0, 16, 16)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/bg.mp3")
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2),
		rl.NewVector2(plrDest.X-(plrDest.Width/2), plrDest.Y-(plrDest.Height/2)),
		0.0, 2)
	loadMap("one.map")
}

func loadMap(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	tilesAndSrc := bytes.Split(f, []byte("\n\n"))
	tiles, src := tilesAndSrc[0], tilesAndSrc[1]

	for _, v := range strings.Fields(string(tiles)) {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if mapW == 0 {
			mapW = n
			continue
		}
		if mapH == 0 {
			mapH = n
			continue
		}
		tileMap = append(tileMap, n)
	}

	for _, v := range bytes.Fields(src) {
		srcMap = append(srcMap, v...)
	}
}

func drawScene() {
	w, h := tileDest.Width, tileDest.Height
	for i, t := range tileMap {
		tileDest.X = tileDest.Width * float32(i%mapW)
		tileDest.Y = tileDest.Height * float32(i/mapW)

		switch srcMap[i] {
		case 'g':
			tex = grassSprite
		case 'l':
			tex = hillSprite
		case 'f':
			tex = fenceSprite
			tileSrc.X, tileSrc.Y, w, h = 0, 0, tileDest.Width, tileDest.Height
			rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(w, h), 0, rl.White)
		case 'h':
			tex = houseSprite
			tileSrc.X, tileSrc.Y, w, h = 0, 0, tileDest.Width, tileDest.Height
			rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(w, h), 0, rl.White)
		case 'w':
			tex = waterSprite
		case 't':
			tex = tilledSprite
		default:
			log.Fatal("invalid source mapping letter:", string(srcMap[i]))
		}
		tileSrc.X = tileSrc.Width * float32((t-1)%int(tex.Width/int32(tileSrc.Width)))
		tileSrc.Y = tileSrc.Height * float32((t-1)/int(tex.Width/int32(tileSrc.Width)))
		rl.DrawTexturePro(
			tex,
			tileSrc,
			tileDest,
			rl.NewVector2(w, h),
			0,
			rl.White)
	}
	rl.DrawTexturePro(plrSprite, plrSrc, plrDest,
		rl.NewVector2(plrDest.Width, plrDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		plrMoving = true
		plrDir = 1
		plrUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		plrMoving = true
		plrDir = 0
		plrDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		plrMoving = true
		plrDir = 2
		plrLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		plrMoving = true
		plrDir = 3
		plrRight = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()

	plrSrc.X = 0
	if plrMoving {
		if plrUp {
			plrDest.Y -= plrSpeed
		}
		if plrDown {
			plrDest.Y += plrSpeed
		}
		if plrLeft {
			plrDest.X -= plrSpeed
		}
		if plrRight {
			plrDest.X += plrSpeed
		}
		if frameCount%8 == 1 {
			plrFrame++
		}
		plrSrc.X = plrSrc.Width * float32(plrFrame)
	}

	frameCount++
	if plrFrame > 3 {
		plrFrame = 0
	}

	plrSrc.Y = plrSrc.Height * float32(plrDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(plrDest.X-(plrDest.Width/2), plrDest.Y-(plrDest.Height/2))

	plrMoving = false
	plrUp, plrDown, plrLeft, plrRight = false, false, false, false
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgClr)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(hillSprite)
	rl.UnloadTexture(fenceSprite)
	rl.UnloadTexture(houseSprite)
	rl.UnloadTexture(waterSprite)
	rl.UnloadTexture(tilledSprite)

	rl.UnloadTexture(plrSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
