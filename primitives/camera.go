package primitives

import (
  "math"
)

type Camera struct {
  lowerLeft, horizontal, vertical, origin, u, v, w Vector
  lensRadius                                       float64

}

func NewCameraCentered(vfov, aspect, aperture float64) Camera {
  lookFrom := Vector{0, 0, 0}
  lookAt := Vector{0, 0, -1}
  focusDist := lookFrom.Sub(lookAt).Length()
  return NewCamera(lookFrom, lookAt , Vector{0, 1, 0}, vfov, aspect, aperture, focusDist)
}


func NewCamera(lookFrom, lookAt, vUp Vector, vFov, aspect, aperture, focusDist float64) Camera {
  c := Camera{}

  c.origin = lookFrom
  c.lensRadius = aperture / 2

  theta := vFov * math.Pi / 180
  halfHeight := math.Tan(theta / 2)
  halfWidth := aspect * halfHeight

  w := lookFrom.Sub(lookAt).Normalize()
  u := vUp.Cross(w).Normalize()
  v := w.Cross(u)

  x := u.MultiplyScalar(halfWidth * focusDist)
  y := v.MultiplyScalar(halfHeight * focusDist)

  c.lowerLeft = c.origin.Sub(x).Sub(y).Sub(w.MultiplyScalar(focusDist))
  c.horizontal = x.MultiplyScalar(2)
  c.vertical = y.MultiplyScalar(2)

  c.w = w
  c.u = u
  c.v = v

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