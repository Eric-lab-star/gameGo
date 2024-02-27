package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Msg        string
	koreanFont rl.Font
)

func init() {

	fmt.Println("Compiling...")
	screenWidth := int32(800)
	screenHeight := int32(1000)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	rl.SetTargetFPS(60)

}

func main() {

	Msg = "안녕하세요"
	koreanFont = rl.LoadFontEx("./fonts/Kart_gothic_medium.ttf", 96, []rune(Msg))
	counter := 0

	for !rl.WindowShouldClose() {
		if counter < 150 {
			counter++
		}

		if ((counter / 10) % 3) == 0 {

			draw(counter / 10)
		}

	}
	rl.UnloadFont(koreanFont)
	rl.CloseWindow()
}

func draw(frame int) {

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.DrawTextEx(koreanFont, Msg[0:frame], rl.Vector2{X: 0, Y: float32(0)}, 96, 0, rl.Black)
	rl.EndDrawing()
}
