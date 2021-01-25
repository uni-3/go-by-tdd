package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	W float64
	H float64
}

// mathod receiver func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.W * r.H
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

type Triangle struct {
	B float64
	H float64
}

func (t Triangle) Area() float64 {
	return t.B * t.H * 0.5
}

// Perimeter 長方形の周囲の長さ
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.W + rectangle.H)
}

// Area 長方形の面積
func Area(rectangle Rectangle) float64 {
	return rectangle.W * rectangle.H
}
