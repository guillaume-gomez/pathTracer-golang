package primitives


import (
  "fmt"
)

type Box struct {
  a, b Vector
  Material
}

func NewBox(p1, p2 Vector, m Material ) Box {
  fmt.Printf("\n{%d ; %d}", p1.ToArray(), p2.ToArray())
  return Box{p1, p2, m}
}


// func (box Box) Position() Vector {
//   return box.position
// }

// func (box Box) Size() Vector {
//   return box.size
// }

// func (box Box) Width() float64 {
//   return box.size.X
// }

// func (box Box) Height() float64 {
//   return box.size.Y
// }

// func (box Box) Length() float64 {
//   return box.size.Z
// }

func (box Box) A() Vector {
  return box.a
}

func (box Box) B() Vector {
  return box.b
}



func(box * Box) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: box.Material}
  rec.Normal = Vector{0,0,1}
  tmin := (r.Origin().X - box.a.X) / r.Direction().X
  tmax := (r.Origin().X - box.b.X) / r.Direction().X

  if tmin > tmax {
    temp := tmin
    tmin = tmax
    tmax = temp
    rec.Normal = Vector{0,0,-1}
  }

  tymin := (r.Origin().Y - box.a.Y) / r.Direction().Y
  tymax := (r.Origin().Y - box.b.Y) / r.Direction().Y

  if tymin > tymax {
    temp := tymin
    tymin = tymax
    tymax = temp
  }

  if (tmin > tymax) || (tymin > tmax) {
    return false, rec
  }

  if tymin > tmin {
    tmin = tymin
    //rec.Normal = Vector{0,-1,0}
  }

  if tymax < tmax {
    tmax = tymax
    //rec.Normal = Vector{0,1,0}
  }

  tzmin := (r.Origin().Z - box.a.Z) / r.Direction().Z
  tzmax := (r.Origin().Z - box.b.Z) / r.Direction().Z

  if tzmin > tzmax {
    temp := tzmin
    tzmin = tzmax
    tzmax = temp
  }

  if (tmin > tzmax) || (tzmin > tmax) {
    return false, rec
  }

  if (tzmin > tmin) {
    tmin = tzmin
  }

  if (tzmax < tmax) {
    tmax = tzmax
  }

  if (tmin < tMin) || (tmin > tMax) {
      return false, rec
    }

    rec.T = tmin
    rec.Point = r.Point(tmin)
    return true, rec
 }
