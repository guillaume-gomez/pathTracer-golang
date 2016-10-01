package primitives

type Metal struct {
  C    Vector
  Fuzz float64
}

func (m Metal) Bounce(input Ray, hit HitRecord) (bool, Ray) {
  direction := reflect(input.Direction(), hit.Normal)
  fuzzed := VectorInUnitSphere().MultiplyScalar(m.Fuzz)
  bouncedRay := Ray{hit.Point, direction.Add(fuzzed)}
  bounced := direction.Dot(hit.Normal) > 0
  return bounced, bouncedRay
}


func (m Metal) Color() Vector {
  return m.C
}
