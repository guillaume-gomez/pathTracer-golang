package primitives

type Disk struct {
  Plane
  radius float64
}




func NewDisk(point, normal Vector, m Material, radius float64) Disk {
  return Disk{Plane{point, normal, m}, radius}
}

func(d *Disk) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: d.Material}
  return false, rec
}
