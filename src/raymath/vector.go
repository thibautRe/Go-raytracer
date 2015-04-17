package raymath

import "math"

type Vector struct {
    X, Y, Z float64
}

var VectorNull = Vector{0,0,0}

func NewVector(p1 Point, p2 Point) Vector {
    return Vector{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
}

////// METHODS //////

func (v1 Vector) Plus(v2 Vector) Vector {
    return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector) Minus(v2 Vector) Vector {
    return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v Vector) Mult(f float64) Vector {
    return Vector{v.X*f, v.Y*f, v.Z*f}
}

func (v Vector) GetUnit() Vector {
    a := v.Abs()
    return Vector{v.X/a, v.Y/a, v.Z/a}
}

func (v Vector) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v1 Vector) ScalarProduct(v2 Vector) float64 {
    return v1.X * v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector) VectorialProduct(v2 Vector) Vector {
    return Vector{
        v1.Y * v2.Z - v1.Z * v2.Y,
        v1.Z * v2.X - v1.X * v2.Z,
        v1.X * v2.Y - v1.Y * v2.X,
    }
}