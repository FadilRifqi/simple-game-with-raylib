package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const(
	screenWidth = 1000
	screenHeight = 480
)

var(
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	hillSprite rl.Texture2D
	fenceSprite rl.Texture2D
	houseSprite rl.Texture2D
	waterSprite rl.Texture2D
	tilledSprite rl.Texture2D

	tex rl.Texture2D

	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerMoving bool
	playerDir int
	playerUp, playerDown, playerLeft, playerRight bool
	playerFrame int

	frameCount int

	tileDest rl.Rectangle
	tileSrc rl.Rectangle
	tileMap []int
	srcMap []string
	mapW, mapH int

	playerSpeed float32 = 3

	musicPaused bool
	music rl.Music

	cam rl.Camera2D
)

func drawScene(){
	// rl.DrawTexture(grassSprite, 100, 50, rl.White)

	for i:=0; i < len(tileMap); i++ {
		if tileMap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i % mapW)
			tileDest.Y = tileDest.Height * float32(i / mapW)

			if srcMap[i] == "g" {
				tex = grassSprite
			}
			if srcMap[i] == "h" {
				tex = hillSprite
			}
			if srcMap[i] == "f" {
				tex = fenceSprite
			}
			if srcMap[i] == "r" {
				tex = houseSprite
			}
			if srcMap[i] == "w" {
				tex = waterSprite
			}
			if srcMap[i] == "t" {
				tex = tilledSprite
			}

			tileSrc.X = tileSrc.Width * float32((tileMap[i] - 1) % int(tex.Width / int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32((tileMap[i] - 1) / int(tex.Width / int32(tileSrc.Width)))
			rl.DrawTexturePro(tex, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
		}
	}

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

func loadMap(mapFile string) {
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	removeNewLines := strings.Replace(string(file), "\r\n", " ", -1)
	sliced := strings.Split(removeNewLines, " ")
	mapW = -1
	mapH = -1

	for i := 0; i < len(sliced); i++ {
		s, _ :=  strconv.ParseInt(sliced[i], 10, 64)
		m:= int(s)


		if mapW == -1 {
			mapW = m
		}else if mapH == -1 {
			mapH = m
		}else if i < mapW * mapH+2 {
			tileMap = append(tileMap, m)
		}else {
			srcMap = append(srcMap, sliced[i])
		}
	}
	if len(tileMap) > mapH * mapW {
		tileMap = tileMap[:len(tileMap) - 1]
	}
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
	if grassSprite.ID == 0 {
        fmt.Println("Failed to load grass sprite texture")
        os.Exit(1)
    }

	hillSprite = rl.LoadTexture("resource/Tilesets/Hills.png")
	if hillSprite.ID == 0 {
		fmt.Println("Failed to load dirt sprite texture")
		os.Exit(1)
	}

	fenceSprite = rl.LoadTexture("resource/Tilesets/Fences.png")
	if fenceSprite.ID == 0 {
		fmt.Println("Failed to load fence sprite texture")
		os.Exit(1)
	}

	houseSprite = rl.LoadTexture("resource/Tilesets/Wooden_House_Roof_Tilset.png")
	if houseSprite.ID == 0 {
		fmt.Println("Failed to load house sprite texture")
		os.Exit(1)
	}

	waterSprite = rl.LoadTexture("resource/Tilesets/Water.png")
	if waterSprite.ID == 0 {
		fmt.Println("Failed to load water sprite texture")
		os.Exit(1)
	}

	tilledSprite = rl.LoadTexture("resource/Tilesets/Tilled_Dirt_v2.png")
	if tilledSprite.ID == 0 {
		fmt.Println("Failed to load tilled sprite texture")
		os.Exit(1)
	}

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc = rl.NewRectangle(0, 0, 16, 16)

	playerSprite = rl.LoadTexture("resource/Characters/BasicCharakterSpritesheet.png")
	if playerSprite.ID == 0 {
        fmt.Println("Failed to load player sprite texture")
        os.Exit(1)
    }


	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("resource/song.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
    rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X - (playerDest.Width/2)), float32(playerDest.Y - (playerDest.Height/2))), 0, 1.5 )

	loadMap("one.map")
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadTexture(hillSprite)
	rl.UnloadTexture(fenceSprite)
	rl.UnloadTexture(houseSprite)
	rl.UnloadTexture(waterSprite)
	rl.UnloadTexture(tilledSprite)
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
