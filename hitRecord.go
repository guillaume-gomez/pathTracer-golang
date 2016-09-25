package main

type HitRecord struct {
  T         float64
  P, Normal Vector
}

func buildHitRecord(t float64, ray *Ray, sphere *Sphere) HitRecord {
  return HitRecord{
    T: t,
    P: ray.Point(t),
    Normal: (ray.Point(t).Sub(sphere.Center)).DivideScalar(sphere.Radius) }
}

