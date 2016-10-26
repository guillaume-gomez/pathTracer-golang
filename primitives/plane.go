package primitives

type Plane struct {
  p, normal Vector
  Material
}


func(p Plane) Point() Vector {
  return p.p
}


func(p Plane) Normal() Vector {
  return p.normal
}

func NewPlane(point, normal Vector, m Material) Plane {
  return Plane{point, normal, m}
}

func(p *Plane) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: p.Material}
  // assuming vectors are all normalized
  denom := ray.Direction().Normalize().Dot(p.Normal().Normalize())
  if (denom > 1e-6) {
    pbVector :=  ray.Origin().Sub(p.Point());
    t := (pbVector.Dot(p.Normal())) / denom
    if( t >= 0 && t > tMin && t < tMax) {
      rec.T = t
      rec.Point = ray.Point(t)
      rec.Normal = p.Normal().Invert()
      return true, rec
    }
    return false, rec
  }
  return false, rec
}
