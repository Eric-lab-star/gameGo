package camera

import rl "github.com/gen2brain/raylib-go/raylib"

type View2D struct {
	*rl.Camera2D
}

func New(pos rl.Vector2, screenWidth int32, screenHeight int32) View2D {
	v := View2D{}
	v.Camera2D = &rl.Camera2D{
		Zoom:     1.0,
		Rotation: 0.0,
		Target:   pos,
		Offset:   rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
	}
	return v
}

func (v *View2D) UpdateTarget(pos rl.Vector2) {
	v.Camera2D.Target = pos
}

func (v *View2D) UpdateZoom(zoom float32, opts ...bool) {
	v.Camera2D.Zoom = zoom
}

func (v *View2D) LimitZoom(max float32, min float32) {
	if v.Zoom > max {
		v.Zoom = max
	} else if v.Zoom < min {
		v.Zoom = min
	}
}
