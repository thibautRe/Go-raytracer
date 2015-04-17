package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "image/png"
    "./src/rayobject"
)

func main() {
    var scene rayobject.Scene
    fileContent, err := ioutil.ReadFile("tests/plan.json")
    if err != nil {
        fmt.Println(err)
    }
    err2 := json.Unmarshal(fileContent, &scene);
    if err2 != nil {
        fmt.Println(err2)
    }

    // Save the image
    render := scene.Render()
    file, _ := os.Create("render/test.png")
    png.Encode(file, render)
}