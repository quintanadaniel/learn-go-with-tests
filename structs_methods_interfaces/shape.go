package structsmethodsinterfaces

import (
	"math"
)

type Rectangle struct {
	Width float64
	Heigh float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base  float64
	Heigh float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigh)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigh
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Heigh) / 2
}
