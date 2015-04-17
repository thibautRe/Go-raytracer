package rayobject

import (
    "testing"
    "image/color"
    "../raymath"
    "math/rand"
    "time"
    "fmt"
)

func TestGetBounceRays(t *testing.T) {
    m := Material{"Diffuse", color.RGBA{0,0,0,0}, 0}

    rays := m.GetBounceRays(raymath.Point{0,0,0}, raymath.Vector{0,0,1}, 10)
    if len(rays) != 10 {
        t.Error("len(rays) expected 10, was ", len(rays))
    }
    for i:=0;i<10;i++ {
        if rays[i].Direction.Z < 0 {
            t.Error("Did not expected a vector with Z coordinate < 0")
        }
    }
}