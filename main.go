package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	hello := Header{
		Msg:      "안녕",
		Position: rl.NewVector2(0, 0),
		Spacing:  0,
		Color:    rl.Black,
	}

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

type Header struct {
	Font     rl.Font
	Msg      string
	Position rl.Vector2
	Size     float32
	Spacing  float32
	Color    rl.Color
}

func (h *Header) CenterX(screenWidth float32) {
	h.Position.X = screenWidth/2 - rl.MeasureTextEx(h.Font, h.Msg, h.Size, h.Spacing).X/2
}

func (h *Header) setFont(filename string) {
	h.Font = rl.LoadFontEx(filename, 96, []rune(h.Msg))
}
