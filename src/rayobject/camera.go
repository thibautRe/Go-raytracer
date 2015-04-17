package rayobject

import "../raymath"

type Camera struct {
    Center raymath.Point
    Direction raymath.Vector
    Normal raymath.Vector
    LensFac float64
}

func (c Camera) GetFirstRay(xPercent float64, yPercent float64) raymath.Ray {
    Normal2 := c.Direction.VectorialProduct(c.Normal)
    direction := c.Direction.GetUnit().Plus(c.Normal.GetUnit().Mult(c.LensFac*(1- 2*yPercent))).Plus(Normal2.GetUnit().Mult(c.LensFac*(2*xPercent-1)))
    return raymath.Ray{c.Center, direction}
}