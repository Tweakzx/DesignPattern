package decorator

type IDraw interface {
	Draw() string
}

type Square struct{}

func (s *Square) Draw() string {
	return "This is a square"
}

type ColorSquare struct {
	Square IDraw
	color  string
}

func NewColorSquare(s IDraw, c string) *ColorSquare {
	return &ColorSquare{Square: s, color: c}
}

func (c ColorSquare) Draw() string {
	return c.Square.Draw() + ", color is " + c.color
}
