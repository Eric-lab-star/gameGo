package main

// ok cmd/main.go file

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	msg := "안녕"

	// TTF Font loading with custom generation parameters
	font := rl.LoadFontEx("fonts/Kart_gothic_medium.ttf", 96, []rune(msg))

	fontSize := font.BaseSize
	fontPosition := rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight / 2)}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextEx(font, msg, fontPosition, float32(fontSize), 0, rl.Black)
		rl.EndDrawing()
	}
	rl.UnloadFont(font) // Font unloading
	rl.CloseWindow()
}
