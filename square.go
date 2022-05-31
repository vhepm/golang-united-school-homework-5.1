package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	var x, y int

	x = s.start.x + int(s.a)
	y = s.start.y - int(s.a)

	return Point{x: x, y: y}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
