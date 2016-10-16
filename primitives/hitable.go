package primitives

type HitRecord struct {
  T         float64
  Point, Normal Vector
  Material
}

func buildHitRecordFromSphere(t float64, ray Ray, sphere *Sphere) HitRecord {
  return HitRecord{
    T: t,
    Point: ray.Point(t),
    Normal: (ray.Point(t).Sub(sphere.Center())).DivideScalar(sphere.Radius()),
    Material: sphere.Material,
  }
}

//Work in progress
func buildHitRecordFromBox(t float64, ray Ray, box Box) HitRecord {
  return HitRecord{
    T: t,
    Point: Ray.Point(t),
    Normal: Vector{},
    Material: box.Material,
  }
}

type Hitable interface {
  Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord)
}
