package raymath

import "image/color"

// Multiply 2 colors
func MultiplyColors(c1, c2 color.RGBA) color.RGBA {
    n := 255*255
    return color.RGBA{byte(int(c1.R*c2.R)/n), byte(int(c1.G*c2.G)/n), byte(int(c1.B*c2.B)/n), byte(int(c1.A*c2.A)/n)}
}

// 0 <= f <= 1
func MultiplyColor(c color.RGBA, f float64) color.RGBA {
    return color.RGBA{byte(float64(c.R)*f), byte(float64(c.G)*f), byte(float64(c.B)*f), 255}
}