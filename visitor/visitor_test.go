package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress_Visit(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{"Circle", &Circle{Radius: 3}, 28.274333882308138},
		{"Rectangle", &Rectangle{Width: 3, Height: 4}, 12},
		{"Square", &Triangle{Side1: 3, Side2: 4, Side3: 5}, 6},
	}

	visitor := &AreaCalculator{}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.shape.Accept(visitor)
			assert.Equal(t, tt.area, visitor.Areas[i])
		})
	}
}
