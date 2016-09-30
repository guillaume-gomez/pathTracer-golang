package primatives

type HitRecord struct {
  T         float64
  P, Normal Vector
  Material
}

func buildHitRecord(t float64, ray Ray, sphere *Sphere) HitRecord {
  return HitRecord{
    T: t,
    P: ray.Point(t),
    Normal: (ray.Point(t).Sub(sphere.Center)).DivideScalar(sphere.Radius),
    Material: sphere.Material,
  }
}

type Hitable interface {
  Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord)
}
