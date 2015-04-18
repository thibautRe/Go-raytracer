package raymath

import (
    "testing"
    "image/color"    
)

func TestMixColors(t *testing.T) {
    red := color.RGBA{255, 0, 0, 255}
    green := color.RGBA{0, 255, 0, 255}
    colors := make([]color.RGBA, 2)
    colors[0] = red
    colors[1] = green
    mix := MixColors(colors)
}