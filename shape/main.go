package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{
		base:   3,
		height: 3,
	}
	sq := square{
		sideLength: 2,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Printf("The area of shape is %.2f\n", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
