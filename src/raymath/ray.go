package raymath

type Ray struct {
    Origin Point
    Direction Vector
}

// Returns the distance from a Ray to a Point
func (r Ray) GetDistanceToPoint(p Point) float64 {
    vec := NewVector(r.Origin, p)
    if r.Direction.ScalarProduct(vec) < 0 {
        return float64(1e8)
    }
    return vec.VectorialProduct(r.Direction).Abs()/r.Direction.Abs()
}