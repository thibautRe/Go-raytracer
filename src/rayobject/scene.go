package rayobject

import (
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
            im.Set(x, y, scene.BackgroundColor)

            // Compute First Ray
            scene.Camera.GetRay(float64(x - bounds.Min.X)/float64(bounds.Max.X - bounds.Min.X), float64(y - bounds.Min.Y)/float64(bounds.Max.Y - bounds.Min.Y))
        }
    }

    return im
}