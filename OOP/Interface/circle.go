package main

import "math"

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
