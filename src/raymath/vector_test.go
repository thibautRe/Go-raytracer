package raymath

import (
    "testing"
)

func TestVectorialProduct(t *testing.T) {
    v1 := Vector{1,0,0}
    v2 := Vector{0,1,0}
    v3 := v1.VectorialProduct(v2)
    if v3.X != 0 || v3.Y != 0 || v3.Z != 1 {
        t.Error("Expected {0,0,1}, got ", v3)
    }
}