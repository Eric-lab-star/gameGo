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

	title := text.New("안녕하세요")
	title.CenterX(float32(screenWidth))

	for !rl.WindowShouldClose() {
		title.Size += (rl.GetMouseWheelMove() * 4.0)
		title.MoveWithArrow(0, float32(screenWidth), 0, float32(screenHeight))

		handleFileDrop(&title)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		title.Draw()
		rl.EndDrawing()
	}
	rl.UnloadFont(title.Font)
	rl.CloseWindow()
}

func handleFileDrop(text *text.Text) {

	if rl.IsFileDropped() {
		droppedFiles := []string{}
		droppedFiles = rl.LoadDroppedFiles()
		count := len(droppedFiles)

		if count == 1 {
			rl.UnloadFont(text.Font)
			text.SetFont(droppedFiles[0])
			rl.UnloadDroppedFiles()
		}
	}
}
