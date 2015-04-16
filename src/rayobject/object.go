package rayobject

import (
    "image/color"
    "../raymath"
)

type Object struct {
    Type string
    Size float64
    Center raymath.Point
    Color color.RGBA
}