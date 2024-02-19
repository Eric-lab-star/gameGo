package main

import (
	"fmt"

	text "github.com/eric-lab-star/game/text"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Compiling...")
	screenWidth := int32(800)
	screenHeight := int32(1000)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	rl.SetTargetFPS(60)

	krMsg := "안녕하세요"
	koreanFont := rl.LoadFontEx("./fonts/Kart_gothic_medium.ttf", 96, []rune(krMsg))
	fmt.Println("msg lenght", +len(krMsg))
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for i, h := 0, 96; i < 6; i, h = i+1, h+96 {

			rl.DrawTextEx(koreanFont, krMsg[0:i], rl.Vector2{X: 0, Y: float32(h)}, 96, 0, rl.Black)
		}

		rl.EndDrawing()
	}

	rl.UnloadFont(koreanFont)
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
