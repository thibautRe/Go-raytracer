package raymath


// Point structure
type Point struct {
    X, Y, Z float64
}

// Origin
var Orig = Point{0,0,0}

///// METHODS //////

func (p1 Point) Distance(p2 Point) float64 {
    return NewVector(p1, p2).Abs()
}

