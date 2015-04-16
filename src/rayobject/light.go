package rayobject

import (
    "image/color"
    "../raymath"
)

type Light struct {
    Type string
    Color color.RGBA
    Center raymath.Point
}