package main

import rl "github.com/gen2brain/raylib-go/raylib"

const(
	screenWidth = 1000
	screenHeight = 480
)

var(
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle

	playerSpeed float32 = 3
)

func drawScene(){
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bgColor)

	drawScene()
	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
}

func initialize() {
	rl.InitWindow(screenWidth, screenHeight, "Sproud Lands")
	rl.SetWindowState(rl.FlagWindowResizable);
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	grassSprite = rl.LoadTexture("resource/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("resource/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
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
