package building

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MaxBuildings int = 100
)

type Building struct {
	Shape rl.Rectangle
	Color rl.Color
}

type Buildings []Building

func CreateBuildings(screenHeight int32, land float32) Buildings {

	buildings := make([]Building, MaxBuildings)

	spacing := float32(0)

	for i := 0; i < MaxBuildings; i++ {
		buildings[i].Shape = rl.Rectangle{
			Width:  float32(rl.GetRandomValue(50, 200)),
			Height: float32(rl.GetRandomValue(100, 800)),
			Y:      float32(screenHeight) - land,
			X:      -6000 + spacing,
		}
		buildings[i].Shape.Y -= buildings[i].Shape.Height
		spacing += buildings[i].Shape.Width

		buildings[i].Color = rl.NewColor(
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(255),
		)
	}
	return buildings
}

func (bs Buildings) Draw() {
	for i := 0; i < MaxBuildings; i++ {
		rl.DrawRectangleRec(bs[i].Shape, bs[i].Color)
	}
}
