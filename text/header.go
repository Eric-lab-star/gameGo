package text

import rl "github.com/gen2brain/raylib-go/raylib"

func New(msg string) Text {
	t := Text{
		Msg:      msg,
		Position: rl.NewVector2(0, 0),
		Spacing:  0,
		Color:    rl.Black,
	}
	t.SetFont("fonts/Kart_gothic_medium.ttf")
	t.Size = float32(t.Font.BaseSize)
	return t
}

type Text struct {
	Font     rl.Font
	Msg      string
	Position rl.Vector2
	Size     float32
	Spacing  float32
	Color    rl.Color
}

func (t *Text) MoveWithArrow(minX float32, maxX float32, minY float32, maxY float32) {
	switch {
	case rl.IsKeyDown(rl.KeyLeft):
		if t.Position.X > minX+10 {
			t.Position.X -= 10
		}
	case rl.IsKeyDown(rl.KeyRight):
		if t.Position.X < maxX-t.Len() {
			t.Position.X += 10
		}
	case rl.IsKeyDown(rl.KeyUp):
		if t.Position.Y > minY+10 {
			t.Position.Y -= 10
		}
	case rl.IsKeyDown(rl.KeyDown):
		if t.Position.Y < maxY-t.Heigth() {
			t.Position.Y += 10
		}
	}
}

func (t *Text) Len() float32 {
	return rl.MeasureTextEx(t.Font, t.Msg, t.Size, t.Spacing).X
}

func (t *Text) Heigth() float32 {
	return rl.MeasureTextEx(t.Font, t.Msg, t.Size, t.Spacing).Y
}

func (t *Text) CenterX(screenWidth float32) {
	t.Position.X = screenWidth/2 - rl.MeasureTextEx(t.Font, t.Msg, t.Size, t.Spacing).X/2
}

func (t *Text) SetFont(filename string) {
	t.Font = rl.LoadFontEx(filename, 96, []rune(t.Msg))
}

func (t *Text) Draw() {
	rl.DrawTextEx(t.Font, t.Msg, t.Position, t.Size, t.Spacing, t.Color)
}
