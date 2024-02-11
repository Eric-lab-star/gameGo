package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	msg := "안녕"

	// NOTE: Textures/Fonts MUST be loaded after Window initialization (OpenGL context is required)

	// TTF Font loading with custom generation parameters
	font := rl.LoadFontEx("fonts/Kart_gothic_medium.ttf", 96, []rune(msg))

	fontSize := font.BaseSize
	fontPosition := rl.NewVector2(40, float32(screenHeight)/2+50)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		fontSize += int32(rl.GetMouseWheelMove() * 4.0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextEx(font, msg, fontPosition, float32(fontSize), 0, rl.Black)
		rl.EndDrawing()
	}

	rl.UnloadFont(font) // Font unloading
	rl.CloseWindow()
}
