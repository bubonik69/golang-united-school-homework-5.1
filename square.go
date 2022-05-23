package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func newSquare(x Point,l uint) (Square){
	return Square{x,l}
}

func (s Square) End() Point {
	// implement me
	return Point{s.start.x + int(s.a), s.start.y+int(s.a)}
}



func (s Square) Area() uint {
	// implement me
	return s.a*s.a
}

func (s Square) Perimeter() uint {
	// implement me
	return s.a*4
}