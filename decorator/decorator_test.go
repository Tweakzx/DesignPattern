package decorator

import (
	"testing"
)

func TestDecorator_Draw(t *testing.T) {
	tests := []struct {
		color string
		want  string
	}{
		{"red", "This is a square, color is red"},
		{"blue", "This is a square, color is blue"},
		{"green", "This is a square, color is green"},
	}

	for _, tt := range tests {
		t.Run(tt.color, func(t *testing.T) {
			square := &Square{}
			decorator := NewColorSquare(square, tt.color)
			if got := decorator.Draw(); got != tt.want {
				t.Errorf("Decorator.Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
