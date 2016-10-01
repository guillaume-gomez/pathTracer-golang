package primitives

import (
  "math"
)

type Camera struct {
  lowerLeft, horizontal, vertical, origin Vector
}

func NewCamera(vfov, aspect float64) Camera {
  return NewCameraWithPosition(Vector{0, 0, 0}, Vector{0, 0, -1}, Vector{0, 1, 0}, vfov, aspect)
}


func NewCameraWithPosition(lookFrom, lookAt, vUp Vector, vfov, aspect float64) Camera {

  theta := vfov * math.Pi / 180
  halfHeight := math.Tan(theta / 2)
  halfWidth := aspect * halfHeight

  c := Camera{}

  c.origin = lookFrom

  w := lookFrom.Sub(lookAt).Normalize()
  u := vUp.Cross(w).Normalize()
  v := w.Cross(u)

  c.lowerLeft = c.origin.Sub(u.MultiplyScalar(halfWidth)).Sub(v.MultiplyScalar(halfHeight)).Sub(w)
  c.horizontal = u.MultiplyScalar(2 * halfWidth)
  c.vertical = v.MultiplyScalar(2 * halfHeight)

  return c
}

func (c *Camera) RayAt(u float64, v float64) Ray {
  position := c.position(u, v)
  direction := c.direction(position)

  return Ray{c.origin, direction}
}

func (c *Camera) position(u float64, v float64) Vector {
  horizontal := c.horizontal.MultiplyScalar(u)
  vertical := c.vertical.MultiplyScalar(v)

  return horizontal.Add(vertical)
}

func(c* Camera) MoveTo(newPosition Vector) {
  c.origin =  newPosition
}

func (c *Camera) direction(position Vector) Vector {
  return c.lowerLeft.Add(position)
}