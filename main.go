package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	maxBuildings int   = 100
	screenWidth  int32 = 800
	screenHeight int32 = 450
	player       Player2D
	buildings    []rl.Rectangle = make([]rl.Rectangle, maxBuildings)
	buildColors  []rl.Color     = make([]rl.Color, maxBuildings)
	spacing      float32        = 0
	land         float32        = 130
	camera       View2D
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "2d camera")
	createBuildings()
	player.New()

	rl.SetTargetFPS(60)
}

type View2D struct {
	*rl.Camera2D
}

func (v *View2D) New() {
	v.Camera2D = &rl.Camera2D{
		Zoom:     1.0,
		Rotation: 0.0,
		Target:   player.Position(),
		Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
	}
}

func (v *View2D) Update(target rl.Vector2) {
	v.Target = target
}

type Player2D struct {
	*rl.Rectangle
}

func (p *Player2D) New() {
	p.Rectangle = &rl.Rectangle{
		X:      400,
		Y:      280,
		Width:  40,
		Height: 40,
	}
}

func (p *Player2D) Position() rl.Vector2 {
	return rl.NewVector2(p.X/2, p.Y/2)

}

func (p *Player2D) Draw() {
	rl.DrawRectangleRec(*p.Rectangle, rl.Red)
}

func createBuildings() {
	for i := 0; i < maxBuildings; i++ {
		r := rl.Rectangle{
			Width:  float32(rl.GetRandomValue(50, 200)),
			Height: float32(rl.GetRandomValue(100, 800)),
			Y:      float32(screenHeight) - land,
			X:      -6000 + spacing,
		}
		r.Y -= r.Height
		spacing += r.Width

		c := rl.NewColor(
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(255),
		)
		buildings[i] = r
		buildColors[i] = c
	}
}

func main() {
	gameLoop()
	rl.CloseWindow()
}

func gameLoop() {
	for !rl.WindowShouldClose() {
		gameInput()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBuildings()
		player.Draw()

		rl.EndDrawing()
	}
}

func drawBuildings() {
	for i := 0; i < maxBuildings; i++ {
		rl.DrawRectangleRec(buildings[i], buildColors[i])
	}
}

func gameInput() {
	switch {
	case rl.IsKeyDown(rl.KeyRight):
		player.X += 2
	case rl.IsKeyDown(rl.KeyLeft):
		player.X -= 2
	}
}
