package primatives

type World struct {
  elements []Hitable
}

func (w *World) Add(h Hitable) {
  w.elements = append(w.elements, h)
}

func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
  hitAnything := false
  closest := tMax
  record := HitRecord{}
  for _, element := range w.elements {
    hit, tempRecord := element.Hit(r, tMin, closest)
    if hit {
      hitAnything = true
      closest = tempRecord.T
      record = tempRecord
    }
  }
  return hitAnything, record
}