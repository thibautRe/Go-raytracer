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
    Fac float64
}

type MaterialLightray struct {
    Lightray raymath.Lightray
    Material Material
}

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func (m Material) GetBounceRays(point raymath.Point, normal raymath.Vector, incidentRay raymath.Ray, number int) ([]raymath.Ray, int) {
    switch m.Type {
    case "Diffuse":
        bounceRays := make([]raymath.Ray, number)
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
        return bounceRays, 1
    case "Emit":
        return make([]raymath.Ray, 0), 1
    case "Transparent":
        bounceRays := make([]raymath.Ray, 1)
        bounceRays[0].Origin = point
        bounceRays[0].Direction = incidentRay.Direction
        return bounceRays, 0
    }
    bounceRays := make([]raymath.Ray, 0)
    return bounceRays, 1
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
    case "Transparent":
        // Only 1 lightray
        lightray.Color = lightrays[0].Color
        lightray.Power = lightrays[0].Power
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
    case "Transparent":
        l.Color = color.RGBA{0,0,0,0}
        l.Power = 0
    }
    return l
}

func MixMaterialLightray(matlrs []MaterialLightray) raymath.Lightray {
    var meanR, meanG, meanB, meanA float64
    var sumFac float64
    var l raymath.Lightray
    var power float64
    for _, matlr := range matlrs {
        sumFac += matlr.Material.Fac
        meanR += float64(matlr.Lightray.Color.R)*matlr.Material.Fac
        meanG += float64(matlr.Lightray.Color.G)*matlr.Material.Fac
        meanB += float64(matlr.Lightray.Color.B)*matlr.Material.Fac
        meanA += float64(matlr.Lightray.Color.A)*matlr.Material.Fac
        power += matlr.Lightray.Power*matlr.Material.Fac
    }
    meanR /= sumFac
    meanG /= sumFac
    meanB /= sumFac
    meanA /= sumFac
    power /= sumFac
    l.Color = color.RGBA{uint8(meanR), uint8(meanG), uint8(meanB), uint8(meanA)}
    l.Power = power
    return l
}