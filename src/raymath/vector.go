package raymath

import "math"

type Vector struct {
    X, Y, Z float64
}

func NewVector(p1 Point, p2 Point) Vector {
    return Vector{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

////// METHODS //////


func (v1 Vector) Minus(v2 Vector) Vector {
    return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vector) GetUnit() Vector {
    a := v.Abs()
    return Vector{v.X/a, v.Y/a, v.Z/a}
}

func (v Vector) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}