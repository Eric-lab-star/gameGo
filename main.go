package main

import (
	"fmt"

	text "github.com/eric-lab-star/game/text"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Compiling...")
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	rl.SetTargetFPS(60)

	title := text.CreateHeader("안녕하세요")
	title.CenterX(float32(screenWidth))

	for !rl.WindowShouldClose() {
		title.Size += (rl.GetMouseWheelMove() * 4.0)
		title.Move(0, float32(screenWidth))

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		title.Draw()
		rl.EndDrawing()
	}
	rl.UnloadFont(title.Font)
	rl.CloseWindow()
}
