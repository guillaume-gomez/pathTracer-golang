package primitives

import (
  "math"
)

type Disk struct {
  Plane
  radius float64
}




func NewDisk(point, normal Vector, m Material, radius float64) Disk {
  return Disk{Plane{point, normal, m}, radius}
}

func(d *Disk) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: d.Material}
  hit, t := d.HitAndGetT(ray, tMin, tMax)
  if hit {
    intersection := ray.Point(t)
    distance := d.p.Sub(intersection)
    squareDistance := distance.Dot(distance)
    if math.Sqrt(squareDistance) <= d.radius {
      rec.T = t
      rec.Point = intersection
      rec.Normal = d.Normal().Invert()
      return true, rec
    }
  }
  return false, rec
}
