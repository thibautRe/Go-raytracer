package raymath

import (
    "image/color"
)

type Ray struct {
    Origin Point
    Direction Vector
}

type Lightray struct {
    Ray Ray
    Color color.RGBA
    Power float64
}

// Returns the distance from a Ray to a Point
func (r Ray) GetDistanceToPoint(p Point) float64 {
    vec := NewVector(r.Origin, p)
    if r.Direction.ScalarProduct(vec) < 0 {
        return float64(1e8)
    }
    return vec.VectorialProduct(r.Direction).Abs()/r.Direction.Abs()
}