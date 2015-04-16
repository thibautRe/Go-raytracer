package rayobject

import (
    "math"
    "../raymath"
)

type Sphere Object

func (s Sphere) Intersection(r raymath.Ray) (float64, raymath.Point, raymath.Vector) {
    distance := r.GetDistanceToPoint(s.Center)

    // If the ray is too far away from the center of the sphere
    if distance > s.Size || distance <= 0 {
        return 0, raymath.Orig, raymath.VectorNull
    }

    // Get the distance to the intersection from
    // the ray to the sphere
    BA := raymath.NewVector(r.Origin, s.Center)
    scalar := r.Direction.ScalarProduct(BA)
    norm := r.Direction.Abs()
    delta := math.Pow(2*scalar, 2) - 4*norm*norm*(math.Pow(BA.Abs(), 2) - s.Size*s.Size)
    a := (2*scalar - math.Sqrt(delta))/(2*norm)
    dist := a*math.Sqrt(norm)

    impactPoint := raymath.Point{
        r.Origin.X + a*r.Direction.X,
        r.Origin.Y + a*r.Direction.Y,
        r.Origin.Z + a*r.Direction.Z,
    }
    normal := raymath.Vector{
        impactPoint.X - s.Center.X,
        impactPoint.Y - s.Center.Y,
        impactPoint.Z - s.Center.Z,
    }

    return dist, impactPoint, normal
}