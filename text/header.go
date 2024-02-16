package text

import rl "github.com/gen2brain/raylib-go/raylib"

func CreateHeader(msg string) Header {
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

func (h *Header) Move(minX float32, maxX float32) {
	switch {
	case rl.IsKeyDown(rl.KeyLeft):
		if h.Position.X > minX {
			h.Position.X -= 10
		}
	case rl.IsKeyDown(rl.KeyRight):
		if h.Position.X < maxX-h.Len() {
			h.Position.X += 10
		}
	case rl.IsKeyDown(rl.KeyDown):
	case rl.IsKeyDown(rl.KeyDown):

	}
}

func (h *Header) Len() float32 {
	return rl.MeasureTextEx(h.Font, h.Msg, h.Size, h.Spacing).X
}

func (h *Header) Heigth() float32 {
	return rl.MeasureTextEx(h.Font, h.Msg, h.Size, h.Spacing).Y
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
