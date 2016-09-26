package primatives

import (
  "math"
)

type Sphere struct {
  Center Vector
  Radius float64
}

func(s *Sphere) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
  // a*a + 2ab + c
  oc := r.Origin.Sub(s.Center)
  a := r.Direction.Dot(r.Direction)
  b := 2 * oc.Dot(r.Direction)
  c := oc.Dot(oc) - s.Radius * s.Radius
  discriminant := b*b - 4*a*c

  rec := HitRecord{}
  // Two solutions
  if discriminant > 0.0 {
    twoA := 2 * a
    sqrtDiscriminant := math.Sqrt(discriminant)
    t := (-b - sqrtDiscriminant) / twoA
    // return the first solution
    if t < tMax && t > tMin {
      rec = buildHitRecord(t, r, s);
      return true, rec
    }
    //return the second solution if t > tMax or t < tMin
    t = (-b + sqrtDiscriminant) / twoA
    if t < tMax && t > tMin {
      rec = buildHitRecord(t, r, s);
      return true, rec
    }
  // Only one solution
  } else if discriminant == 0.0 {
    t := -b / (2* a)
    if t < tMax && t > tMin {
      rec = buildHitRecord(t, r, s);
      return true, rec
    }
  }
  return false, rec
}
