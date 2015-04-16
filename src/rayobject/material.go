package rayobject

import (
    "image/color"
    "math/rand"
    "../raymath"
)

type Material struct {
    Type string
    Color color.RGBA
}

func (m Material) GetBounceRays(point raymath.Point, normal raymath.Vector, number int) []raymath.Ray {
    bounceRays := make([]raymath.Ray, number)
    switch m.Type {
    case "Diffuse":
        for i := 0; i < number; {
            tempVector := raymath.Vector{
                rand.Float64()-0.5,
                rand.Float64()-0.5,
                rand.Float64()-0.5,
            }
            if tempVector.ScalarProduct(normal) > 0 {
                bounceRays[i] = raymath.Ray{point, tempVector}
                i++
            }
        }
    }
    return bounceRays
}