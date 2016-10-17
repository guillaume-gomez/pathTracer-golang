package primitives

type Box struct {
  position, size, a, b Vector
  Material
}

func (box Box) Position() Vector {
  return box.position
}

func (box Box) Size() Vector {
  return box.size
}

func (box Box) Width() float64 {
  return box.size.X
}

func (box Box) Height() float64 {
  return box.size.Y
}

func (box Box) Length() float64 {
  return box.size.Z
}

func (box Box) A() Vector {
  return box.a
}

func (box Box) B() Vector {
  return box.b
}


 func(box * Box) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: box.Material}

  txmin := (box.A().X - r.Origin().X) / r.Direction().X
  txmax := (box.B().X - r.Origin().X) / r.Direction().X

  if (txmin > txmax) {
    swap := txmin
    txmin = txmax
    txmax = swap
  }

  tymin := (box.A().Y - r.Origin().Y) / r.Direction().Y
  tymax := (box.B().Y - r.Origin().Y) / r.Direction().Y

  if (tymin > tymax) {
    swap := tymin
    tymin = tymax
    tymax = swap
  }

  if ((txmin > tymax) || (tymin > txmax)) {
    return false, rec
  }

  if (tymin > txmin) {
    txmin = tymin
  }

  if (tymax < txmax) {
    txmax = tymax
  }

  tzmin := (box.A().Z - r.Origin().Z) / r.Direction().Z
  tzmax := (box.B().Z - r.Origin().Z) / r.Direction().Z

  if (tzmin > tzmax) {
    swap := tzmin
    tzmin = tzmax
    tzmax = swap
  }

  if ((txmin > tzmax) || (tzmin > txmax)) {
    return false, rec
  }

  if (tzmin > txmin) {
    txmin = tzmin
  }

  if (tzmax < txmax) {
    txmax = tzmax
  }
  t := txmin;
  if (t < 0) {
    t = txmax;
    if (t < 0 || t < tMin || t > tMax) {
      return false, rec
    }
  }
  rec = buildHitRecordFromBox(t, r, box)
  return true, rec
 }