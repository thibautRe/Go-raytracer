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
}

func (scene Scene) Render() *image.RGBA {
    // Create the image
    im := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{scene.Width, scene.Height}})

    // Fill the Background
    bounds := im.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            im.Set(x, y, scene.BackgroundColor)
        }
    }

    return im
}