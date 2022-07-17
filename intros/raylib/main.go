package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running     = true
	bkgClr      = rl.NewColor(147, 211, 196, 255)
	grassSprite rl.Texture2D
	plrSprite   rl.Texture2D

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
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sprouts Lands")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	plrSprite = rl.LoadTexture("res/Characters/Basic_Charakter_Spritesheet.png")

	plrSrc = rl.NewRectangle(0, 0, 48, 48)
	plrDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/bg.mp3")
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2),
		rl.NewVector2(plrDest.X-(plrDest.Width/2), plrDest.Y-(plrDest.Height/2)),
		0.0, 1.5)
}

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
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
