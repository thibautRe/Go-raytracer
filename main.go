package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "image"
    "image/color"
    "image/png"
)

type SceneDescriptor struct {
    BackgroundColor color.RGBA
    Height int
    Width int
}

func main() {
    var scene SceneDescriptor
    fileContent, err := ioutil.ReadFile("tests/test.json")
    if err != nil {
        fmt.Println(err)
    }
    err2 := json.Unmarshal(fileContent, &scene);
    if err2 != nil {
        fmt.Println(err2)
    }

    // Create the image
    im := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{scene.Width, scene.Height}})


    // Fill the Background
    bounds := im.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            im.Set(x, y, scene.BackgroundColor)
        }
    }

    // Save the image
    file, _ := os.Create("render/test.png")
    png.Encode(file, im)
}