package square

import "testing"

var cases = map[string]struct {
start        Point
side         uint
expEnd       Point
expArea      uint
expPerimeter uint
}{
"positive start": {
start:        Point{3, 5},
side:         4,
expEnd:       Point{7, 9},
expArea:      16,
expPerimeter: 16,
},
"negative start": {
start:        Point{-432, -678},
side:         128,
expEnd:       Point{-304, -550},
expArea:      16384,
expPerimeter: 512,
},
"zero side": {
start:        Point{312, 534},
side:         0,
expEnd:       Point{312, 534},
expArea:      0,
expPerimeter: 0,
},
}

func TestSquare(t *testing.T) {
	for k,v:=range cases {
		sq:=newSquare(v.start,v.side)
		endPoint:=sq.End()
		if endPoint!= v.expEnd {
			t.Errorf("error in %s case. Need point %v, get %v",k , v.expEnd, endPoint)
		} else{
			t.Logf("Test %s (End point) pass Ok", k)
		}

		area:=sq.Area()
		if area!= v.expArea {
			t.Errorf("error in %s case. Need point %v, get %v",k , v.expArea, area)
		} else{
		t.Logf("Test %s (Area) pass Ok", k)
	}
	perimeter:=sq.Perimeter()
		if perimeter!= v.expPerimeter {
			t.Errorf("error in %s case. Need point %v, get %v",k , v.expPerimeter, perimeter)
		} else{
			t.Logf("Test %s (Perimeter) pass Ok", k)
		}
	}



	}

