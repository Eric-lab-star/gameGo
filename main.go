package main

import (
	bl "github.com/eric-lab-star/game/building"
	pl "github.com/eric-lab-star/game/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  int32   = 800
	screenHeight int32   = 450
	land         float32 = 130
	player       pl.Player2D
	buildings    bl.Buildings
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "2d camera")
	rl.SetTargetFPS(60)
	player = pl.New()
	buildings = bl.CreateBuildings(screenHeight, land)
}

func main() {

	gameLoop()
	rl.CloseWindow()
}

func gameLoop() {
	for !rl.WindowShouldClose() {
		gameUpdates()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		buildings.Draw()
		player.Draw()

		rl.EndDrawing()
	}
}

func gameUpdates() {
	gameInput()
}

func gameInput() {
	switch {
	case rl.IsKeyDown(rl.KeyRight):
		player.X += 2
	case rl.IsKeyDown(rl.KeyLeft):
		player.X -= 2
	}
}
