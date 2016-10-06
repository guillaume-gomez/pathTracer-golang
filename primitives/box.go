package primitives

type Box struct {
  position, size Vector
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


func(b * Box) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  //TODO
}