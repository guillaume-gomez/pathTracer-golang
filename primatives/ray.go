package primatives

import (
  "math"
)


type Ray struct {
  Origin, Direction Vector
}


func (r Ray) Point(t float64) Vector {
  // p(t)  =  A + B * t
  b := r.Direction.MultiplyScalar(t)
  a := r.Origin
  return a.Add(b)
}

func(r Ray) HitSphere(s Sphere) bool {
  // a*a + 2ab + c
  oc := r.Origin.Sub(s.Center)
  a := r.Direction.Dot(r.Direction)
  b := 2 * oc.Dot(r.Direction)
  c := oc.Dot(oc) - s.Radius * s.Radius
  discriminant := b*b - 4*a*c

  return discriminant > 0
}

func(r Ray) Color(sphere Sphere) Vector {
  white := Vector{1.0, 1.0, 1.0}
  blue := Vector{0.5, 0.7, 1.0}
  red := Vector{ 1.0, 0.0, 0.0 }

  hit, _ := sphere.Hit(r, 0.0, math.MaxFloat64)
  if(hit) {
    //a color
    return red
  }
  // make unit vector so y is between -1.0 and 1.0
  unitDirection := r.Direction.Normalize()

  // scale t to be between 0.0 and 1.0
  t := 0.5 * (unitDirection.Y + 1.0)

  // linear blend
  // blended_value = (1 - t) * white + t * blue
  return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func(r Ray) ColorWithSphere() Vector {
  sphere := Sphere{Center: Vector{0, 0, -1}, Radius: 0.5}
  return r.Color(sphere)
}

