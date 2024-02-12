package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	msg := "안녕"
	font := rl.LoadFontEx("fonts/Kart_gothic_medium.ttf", 96, []rune(msg))

	fontSize := font.BaseSize
	fontPosition := rl.NewVector2(float32(screenWidth)/2-rl.MeasureTextEx(font, msg, float32(fontSize), 0.0).X/2, 0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		fontSize += int32(rl.GetMouseWheelMove() * 4.0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextEx(font, msg, fontPosition, float32(fontSize), 0, rl.Black)
		rl.EndDrawing()
	}

	rl.UnloadFont(font) // Font unloading
	rl.CloseWindow()
}
