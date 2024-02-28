package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	maxBuildings int            = 100
	screenWidth  int32          = 800
	screenHeight int32          = 450
	player       rl.Rectangle   = rl.NewRectangle(400, 280, 40, 40)
	buildings    []rl.Rectangle = make([]rl.Rectangle, maxBuildings)
	buildColors  []rl.Color     = make([]rl.Color, maxBuildings)
	spacing      float32        = 0
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "2d camera")
	createBuildings()
	rl.SetTargetFPS(60)
}

func createBuildings() {
	for i := 0; i < maxBuildings; i++ {
		r := rl.Rectangle{
			Width:  float32(rl.GetRandomValue(50, 200)),
			Height: float32(rl.GetRandomValue(100, 800)),
			Y:      float32(screenHeight) - 130,
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
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < maxBuildings; i++ {
			rl.DrawRectangleRec(buildings[i], buildColors[i])
		}

		rl.DrawRectangleRec(player, rl.Red)
		rl.EndDrawing()

	}

}
