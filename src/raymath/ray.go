package raymath

type Ray struct {
    Origin Point
    Direction Vector
}

func (r Ray) GetDistanceToPoint(p Point) float64 {
    return NewVector(r.Origin, p).VectorialProduct(r.Direction).Abs()/r.Direction.Abs()
}