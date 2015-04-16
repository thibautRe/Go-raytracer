package rayobject

import (
    "../raymath"
)

type Face []raymath.Point

// Computes the normal of the face
func (f Face) GetNormal() raymath.Vector {
    v1 := raymath.Vector{
        f[1].X - f[0].X,
        f[1].Y - f[0].Y,
        f[1].Z - f[0].Z,
    }
    v2 := raymath.Vector{
        f[2].X - f[1].X,
        f[2].Y - f[1].Y,
        f[2].Z - f[1].Z,
    }
    return v1.VectorialProduct(v2)
}

// Tests if the face intersects the Ray
// If the face is intersected, the dist
func (f Face) Intersection(r raymath.Ray) (distance float64, collisionPoint raymath.Point, normal raymath.Vector) {
    Fn := f.GetNormal()
    F0 := raymath.NewVector(f[0], r.Origin)
    // If the direction is orthogonal to the normal, 
    // then the ray never touches the plan
    if Fn.ScalarProduct(r.Direction) == 0 {
        return 0, raymath.Orig, raymath.VectorNull
    }

    a := -F0.ScalarProduct(Fn)/r.Direction.ScalarProduct(Fn)

    // If a < 0 then the ray doesn't intersect the face
    if a <= 0 {
        return 0, raymath.Orig, raymath.VectorNull
    }

    impactPoint := raymath.Point{
        r.Origin.X + a*r.Direction.X,
        r.Origin.Y + a*r.Direction.Y,
        r.Origin.Z + a*r.Direction.Z,
    }

    distance = a*r.Direction.Abs()
    return distance, impactPoint, Fn
}