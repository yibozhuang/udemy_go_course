package main

type visitor interface {
	implSquare(*square)
	implCircle(*circle)
	implRectangle(*rectangle)
}

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

var _ shape = (*square)(nil)

func (s *square) getType() string {
	return "Square"
}

func (s *square) accept(v visitor) {
	v.implSquare(s)
}

type circle struct {
	radius int
}

var _ shape = (*circle)(nil)

func (c *circle) getType() string {
	return "Circle"
}

func (c *circle) accept(v visitor) {
	v.implCircle(c)
}

type rectangle struct {
	length int
	width  int
}

var _ shape = (*rectangle)(nil)

func (r *rectangle) getType() string {
	return "Rectangle"
}

func (r *rectangle) accept(v visitor) {
	v.implRectangle(r)
}
