package rayobject

import (
    "../raymath"
)

type Object struct {
    Type string
    Size float64
    Center raymath.Point
    Material Material
    Faces []Face
}

type RenderableObject interface {
    Intersection(r raymath.Ray) (dist float64, col raymath.Point, normal raymath.Vector)
}