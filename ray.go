package main


type Ray struct {
  Origin, Direction Vector
}


func (r Ray) Point(t float64) Vector {
  // p(t)  =  A + B * t
  b := r.Direction.MultiplyScalar(t)
  a := r.Origin
  return a.Add(b)
}

func(r Ray) HitSphere(s sphere) bool {
  // a*a + 2ab + c
  oc := r.Origin.Sub(s.center)
  a := r.Direction.Dot(r.Direction)
  b := 2 * oc.Dot(r.Direction)
  c := oc.Dot(oc) - s.Radius * s.Radius
  discriminant := b*b - 4*a*c

  return discriminant > 0
}

