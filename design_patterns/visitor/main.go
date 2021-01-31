package main

import (
	"fmt"
	"math"
)

type areaCalculator struct {
	area int
}

var _ visitor = (*areaCalculator)(nil)

func (a *areaCalculator) implSquare(s *square) {
	fmt.Printf("Area for square is: %d\n", s.side*s.side)
}

func (a *areaCalculator) implCircle(c *circle) {
	area := math.Pi * float64(c.radius*c.radius)
	fmt.Printf("Area for circle is: %f\n", area)
}

func (a *areaCalculator) implRectangle(r *rectangle) {
	fmt.Printf("Area for rectangle is: %d\n", r.length*r.width)
}

func main() {
	s := &square{
		side: 4,
	}
	c := &circle{
		radius: 2,
	}
	r := &rectangle{
		length: 4,
		width:  5,
	}

	a := &areaCalculator{}
	s.accept(a)
	c.accept(a)
	r.accept(a)
}
