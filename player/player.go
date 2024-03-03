package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player2D struct {
	*rl.Rectangle
	color rl.Color
}

func New() Player2D {
	return Player2D{
		Rectangle: &rl.Rectangle{
			X:      400,
			Y:      280,
			Width:  40,
			Height: 40,
		},
		color: rl.Red,
	}
}

func (p Player2D) Position() rl.Vector2 {
	return rl.NewVector2(p.X/2, p.Y/2)
}

func (p Player2D) Draw() {
	rl.DrawRectangleRec(*p.Rectangle, p.color)
}
