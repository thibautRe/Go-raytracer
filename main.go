package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type SceneDescriptor struct {
    BackgroundColor string
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

    fmt.Println(scene, err);
}