package rayobject

import (
    "../raymath"
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
    for i := range scene.Objects {
        o := scene.Objects[i]
        if o.Type == "Sphere" {
            dist := r.GetDistanceToPoint(o.Center)
            if dist <= o.Size {
                return o.Material.Color
            }
        }
    }

    // Else, return the backgroundcolor
    return scene.BackgroundColor
}