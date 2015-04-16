package rayobject

import "image/color"

type Material struct {
    Color color.RGBA
    EmitColor color.RGBA
    EmitIntensity float64
}