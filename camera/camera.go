package camera

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	screenWidth  int32 = 800
	screenHeight int32 = 450
)

type View2D rl.Camera2D

func (v View2D) New(pos rl.Vector2) {
	v = View2D{
		Zoom:     1.0,
		Rotation: 0.0,
		Target:   pos,
		Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
	}
}
func (v *View2D) Update(target rl.Vector2) {
	v.Target = target
}
