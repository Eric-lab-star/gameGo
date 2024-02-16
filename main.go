package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Compiling...")
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - ttf loading")

	rl.SetTargetFPS(60)

	title := createHeader("안녕하세요")
	title.CenterX(float32(screenWidth))

	for !rl.WindowShouldClose() {
		title.Size += (rl.GetMouseWheelMove() * 4.0)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		title.Draw()
		rl.EndDrawing()
	}
	rl.UnloadFont(title.Font)
	rl.CloseWindow()
}

func createHeader(msg string) Header {
	header := Header{
		Msg:      msg,
		Position: rl.NewVector2(0, 0),
		Spacing:  0,
		Color:    rl.Black,
	}
	header.setFont("fonts/Kart_gothic_medium.ttf")
	header.Size = float32(header.Font.BaseSize)
	return header
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

func (h *Header) Draw() {
	rl.DrawTextEx(h.Font, h.Msg, h.Position, h.Size, h.Spacing, h.Color)
}
