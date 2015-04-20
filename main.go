package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "image/png"
    "./src/rayobject"
    "time"
)

func main() {
    now := time.Now()
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
    fmt.Println("Begin render...", scene.Samples, "samples,", scene.MaxRecursion, "max recursion")
    render := scene.Render()
    fmt.Println("Render finished in", time.Since(now).Seconds(), "seconds")
    file, _ := os.Create("render/test.png")
    png.Encode(file, render)
}