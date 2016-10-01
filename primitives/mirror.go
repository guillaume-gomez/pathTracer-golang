package primitives

type Mirror struct {
  C Vector
}

func (m Mirror) Bounce(input Ray, hit HitRecord) (bool, Ray) {
  direction := reflect(input.Direction(), hit.Normal)
  bouncedRay := Ray{hit.Point, direction}
  bounced := direction.Dot(hit.Normal) > 0
  return bounced, bouncedRay
}

func (m Mirror) Color() Vector {
  return m.C
}
