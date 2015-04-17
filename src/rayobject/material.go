package rayobject

import (
    "image/color"
    "math/rand"
    "../raymath"
    "time"
    //"fmt"
)

type Material struct {
    Type string
    Color color.RGBA
    Intensity float64
}

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
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
    case "Emit":
        return make([]raymath.Ray, 0)
    }
    return bounceRays
}

func (m Material) ComputeLightrayFromLightrays(lightrays []raymath.Lightray) raymath.Lightray {
    var lightray raymath.Lightray
    l := len(lightrays)

    switch m.Type {
    case "Diffuse":
        power := 0.0
        for i := range lightrays {
            power += lightrays[i].Power
        }
        power /= float64(l)
       // fmt.Println(power)
        lightray.Color = raymath.MultiplyColor(m.Color, power)
        lightray.Power = power
    case "Emit":
        lightray.Color = m.Color
        lightray.Power = m.Intensity
    }

    return lightray
}

func (m Material) GetDefaultLightray() raymath.Lightray {
    var l raymath.Lightray

    switch m.Type {
    case "Diffuse":
        l.Color = m.Color
        l.Power = 0.01
    case "Emit":
        l.Color = m.Color
        l.Power = m.Intensity
    }
    return l
}