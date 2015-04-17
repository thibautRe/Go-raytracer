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
    Samples int
    MaxRecursion int
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

            lightray := scene.GetLightray(ray, 0)

            im.Set(x, y, lightray.Color)
        }
    }

    return im
}

func (scene Scene) GetLightray(r raymath.Ray, index int) raymath.Lightray {
    minDist := 1e8
    hasTouched := false
    var dist float64
    var touchedObject Object
    var impactPoint, impactPointObject raymath.Point
    var normal, normalObject raymath.Vector

    // Go through all objects to check intersection
    for i := range scene.Objects {
        o := scene.Objects[i]

        switch o.Type {
        // Test Sphere intersection
        case "Sphere":
            dist, impactPointObject, normalObject = Sphere(o).Intersection(r)
            if dist < minDist && dist > 0 {
                minDist = dist
                hasTouched = true
                touchedObject = o
                impactPoint = impactPointObject
                normal = normalObject
            }
        // Standard triangle mesh
        case "Mesh":
            for i := range o.Faces {
                dist, impactPointObject, normalObject = o.Faces[i].Intersection(r)
                if dist < minDist && dist > 0 {
                    minDist = dist
                    hasTouched = true
                    touchedObject = o
                    impactPoint = impactPointObject
                    normal = normalObject
                }
            }
        }
    }

    if hasTouched {

        if index <= scene.MaxRecursion {
            bounceRays := touchedObject.Material.GetBounceRays(impactPoint, normal, scene.Samples);
            lightRays := make([]raymath.Lightray, len(bounceRays))
            for i := range bounceRays {
                lightRays[i] = scene.GetLightray(bounceRays[i], index + 1)
            }

            lightray := touchedObject.Material.ComputeLightrayFromLightrays(lightRays)
            return lightray
        } else {
            return touchedObject.Material.GetDefaultLightray()
        }
    }

    // Else, return the backgroundcolor
    return raymath.Lightray{
        r,
        scene.BackgroundColor,
        0,
    }
}