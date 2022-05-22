package square

type point struct {
	x int
	y int
}
type square struct{
	a point 	// first point
	b point		// second point
	l int 		// length of side
}

func (s square) perimeter() (area int) {
	return s.l*4
}

func (s square) area() (area int) {
	return s.l*s.l
}
func (s square) returnEndPoint () (b point) {
	return s.b
}

