package main

import (
	bl "github.com/eric-lab-star/game/building"
	ca "github.com/eric-lab-star/game/camera"
	pl "github.com/eric-lab-star/game/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  int32   = 800
	screenHeight int32   = 450
	land         float32 = 130
	player       pl.Player2D
	buildings    bl.Buildings
	view         ca.View2D
)

func init() {

	rl.InitWindow(screenWidth, screenHeight, "2d camera")
	rl.SetTargetFPS(60)
	player = pl.New()
	buildings = bl.CreateBuildings(screenHeight, land)
	view = ca.New(player.Position(), screenWidth, screenHeight)
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
		rl.BeginMode2D(*view.Camera2D)
		rl.DrawRectangle(-6000, 320, 13000, 8000, rl.DarkGray) // ground
		buildings.Draw()
		player.Draw()
		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func gameUpdates() {
	keyboardInputs()
	mouseInputs()
	view.UpdateTarget(player.Position())
}

func mouseInputs() {
	switch {
	case rl.GetMouseWheelMove() != 0:
		view.UpdateZoom(rl.GetMouseWheelMove())
	}
}

func keyboardInputs() {
	switch {
	case rl.IsKeyDown(rl.KeyRight):
		player.X += 2
	case rl.IsKeyDown(rl.KeyLeft):
		player.X -= 2
	case rl.IsKeyDown(rl.KeyA) && view.Rotation >= -30:
		view.Rotation--
	case rl.IsKeyDown(rl.KeyS) && view.Rotation <= 30:
		view.Rotation++
	}

}
