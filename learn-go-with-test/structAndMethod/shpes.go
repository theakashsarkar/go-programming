package structandmethod

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	Redius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.width + rectangle.height)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.width * rectangle.height
// }
