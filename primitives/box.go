package primitives

type Box struct {
  position, size, a, b Vector
  Material
}

func (b Box) Position() Vector {
  return b.position
}

func (b Box) Size() Vector {
  return b.size
}

func (b Box) Width() float64 {
  return b.size.X
}

func (b Box) Height() float64 {
  return b.size.Y
}

func (b Box) Length() float64 {
  return b.size.Z
}

func (b Box) A() Vector {
  return a
}

func (b Box) B() Vector {
  return b
}


 func(b * Box) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  rec := HitRecord{Material: b.Material}

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

  tzmin := (box.a.z - r.Origin().Z) / r.Direction().Z
  tzmax := (box.b.z - r.Origin().Z) / r.Direction().Z

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
  t := tmin;
  if (t < 0) {
    t = tmax;
    if (t < 0 || t < tMin || t > tMax) {
      return false, rec
    }
  }
  rec = buildHitRecordFromBox(t, r, b)
  return true, rec
 }