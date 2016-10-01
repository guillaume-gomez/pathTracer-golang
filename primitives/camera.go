package primitives

import (
  "math"
)

type Camera struct {
  lowerLeft, horizontal, vertical, origin Vector
}

func NewCamera(vfov, aspect float64) Camera {
  return NewCameraWithPosition(0.0, 0.0, 0.0, vfov, aspect)
}

func NewCameraWithPosition(x, y, z , vfov, aspect float64) Camera {

  theta := vfov * math.Pi / 180
  halfHeight := math.Tan(theta / 2)
  halfWidth := aspect * halfHeight

  c := Camera{}

  c.lowerLeft = Vector{-halfWidth, -halfHeight, -1.0}
  c.horizontal = Vector{2 * halfWidth, 0.0, 0.0}
  c.vertical = Vector{0.0, 2 * halfHeight, 0.0}
  c.origin = Vector{x, y, z}

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

func (c *Camera) direction(position Vector) Vector {
  return c.lowerLeft.Add(position)
}