package visitor

import "math"

type Shape interface {
	Accept(visitor Visitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.Visit(r)
}

type Triangle struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

func (t *Triangle) Accept(visitor Visitor) {
	visitor.Visit(t)
}

// 访问者接口
type Visitor interface {
	Visit(shape Shape)
}

// 面积计算访问者
type AreaCalculator struct {
	Areas []float64
}

func (v *AreaCalculator) Visit(shape Shape) {
	switch s := shape.(type) {
	case *Circle:
		v.VisitCircle(s)
	case *Rectangle:
		v.VisitRectangle(s)
	case *Triangle:
		v.VisitTriangle(s)
	default:
		panic("unknown shape type")
	}
}

func (v *AreaCalculator) VisitCircle(circle *Circle) {
	v.Areas = append(v.Areas, math.Pi*circle.Radius*circle.Radius)
}

func (v *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	v.Areas = append(v.Areas, rectangle.Width*rectangle.Height)
}

func (v *AreaCalculator) VisitTriangle(triangle *Triangle) {
	// 使用海伦公式计算三角形的面积
	s := (triangle.Side1 + triangle.Side2 + triangle.Side3) / 2
	v.Areas = append(v.Areas, math.Sqrt(s*(s-triangle.Side1)*(s-triangle.Side2)*(s-triangle.Side3)))
}
