package rayobject

import (
    "../raymath"
)

type Object struct {
    Type string
    Size float64
    Center raymath.Point
    Material Material
}