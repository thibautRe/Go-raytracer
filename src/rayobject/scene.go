package rayobject

import (
    "../raymath"
    "math"
    "image"
    "image/color"
)

type Scene struct {
    BackgroundColor color.RGBA
    Height int
    Width int
    Objects []Object
    Camera Camera
}

func (scene Scene) Render() *image.RGBA {
    // Create the image
    im := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{scene.Width, scene.Height}})

    // Fill the Background
    bounds := im.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            

            // Compute First Ray
            ray := scene.Camera.GetFirstRay(float64(x - bounds.Min.X)/float64(bounds.Max.X - bounds.Min.X), float64(y - bounds.Min.Y)/float64(bounds.Max.Y - bounds.Min.Y))

            im.Set(x, y, scene.GetColor(ray))
        }
    }

    return im
}

func (scene Scene) GetColor(r raymath.Ray) color.RGBA {
    // All objects
    minDist := 1e8
    var touchedObject Object
    hasTouched := false
    for i := range scene.Objects {
        o := scene.Objects[i]
        // Test Sphere intersection
        if o.Type == "Sphere" {
            dist := r.GetDistanceToPoint(o.Center)

            // Condition for the ray to touch
            if dist <= o.Size {
                // Get the distance to the intersection from
                // the ray to the sphere
                BA := raymath.NewVector(r.Origin, o.Center)
                scalar := r.Direction.ScalarProduct(BA)
                norm := r.Direction.Abs()
                delta := math.Pow(2*scalar, 2) - 4*norm*norm*(math.Pow(BA.Abs(), 2) - o.Size*o.Size)
                a := (2*scalar - math.Sqrt(delta))/(2*norm)
                dist := a*math.Sqrt(norm)
                if (dist <= minDist) {
                    minDist = dist
                    hasTouched = true
                    touchedObject = o
                }
            }
        }
    }

    if hasTouched {
        return touchedObject.Material.Color
    }

    // Else, return the backgroundcolor
    return scene.BackgroundColor
}