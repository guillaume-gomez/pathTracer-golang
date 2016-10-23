package primitives

import (
  "math"
)

type Plane struct {
  Material
}


func(p *plane) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: s.Material}
  return false, rec
}