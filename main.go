package main

import rl "github.com/gen2brain/raylib-go/raylib"

const(
	screenWidth = 1280
	screenHeight = 960
)

var(
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerMoving bool
	playerDir int
	playerUp, playerDown, playerLeft, playerRight bool
	playerFrame int

	frameCount int

	playerSpeed float32 = 3

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
)

func drawScene(){
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDir = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}

	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if playerUp {playerDest.Y -= playerSpeed}
		if playerDown {playerDest.Y += playerSpeed}
		if playerLeft {playerDest.X -= playerSpeed}
		if playerRight {playerDest.X += playerSpeed}
		if frameCount % 8 == 1 {playerFrame++}
	}else if frameCount % 45 == 1 {
		playerFrame ++
	}

	frameCount++
	if playerFrame > 3 {playerFrame = 0}
	if !playerMoving && playerFrame > 1 {playerFrame = 0}

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X - (playerDest.Width/2)), float32(playerDest.Y - (playerDest.Height/2)))

	playerMoving = false
	playerUp, playerDown, playerLeft, playerRight = false, false, false, false
}

func initialize() {
	rl.InitWindow(screenWidth, screenHeight, "Sproud Lands")
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	grassSprite = rl.LoadTexture("resource/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("resource/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("resource/song.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
    rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X - (playerDest.Width/2)), float32(playerDest.Y - (playerDest.Height/2))), 0, 1.5 )
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {
	initialize()

	for running {
		input()
		update()
		render()
	}

	quit()
}
