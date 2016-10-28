package primitives

import (
  "math"
  "math/rand"
)

type Vector struct {
  X, Y, Z float64
}

var UnitVector = Vector{1, 1, 1}

func(v Vector) ToArray() [3]float64 {
  return [3]float64{v.X, v.Y, v.Z}
}

func (v Vector) Length() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Dot(o Vector) float64 {
  return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Normalize() Vector {
  l := v.Length()
  return Vector{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector) Add(o Vector) Vector {
  return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector) Mul(o Vector) Vector {
  return Vector{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

func (v Vector) Sub(o Vector) Vector {
  return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}


func (v Vector) AddScalar(t float64) Vector {
  return Vector{v.X + t, v.Y + t, v.Z + t}
}

func (v Vector) SubtractScalar(t float64) Vector {
  return Vector{v.X - t, v.Y - t, v.Z - t}
}

func (v Vector) MultiplyScalar(t float64) Vector {
  return Vector{v.X * t, v.Y * t, v.Z * t}
}

func (v Vector) DivideScalar(t float64) Vector {
    return Vector{v.X / t, v.Y / t, v.Z / t}
}

func(v Vector) SquaredLength() float64 {
  return ( v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func(v Vector) Invert() Vector {
  return Vector{ -v.X, -v.Y, -v.Z }
}

func VectorInUnitSphere() Vector {
  for {
    r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
    p := r.MultiplyScalar(2.0).Sub(UnitVector)
    if p.SquaredLength() >= 1.0 {
      return p
    }
  }
}

func (v Vector) Reflect(o Vector) Vector {
  b := 2 * v.Dot(o)
  return v.Sub(o.MultiplyScalar(b))
}


func (v Vector) Refract(o Vector, n float64) (bool, Vector) {
  uv := v.Normalize()
  uo := o.Normalize()
  dt := uv.Dot(uo)
  discriminant := 1.0 - (n * n * (1 - dt*dt))
  if discriminant > 0 {
    a := uv.Sub(o.MultiplyScalar(dt)).MultiplyScalar(n)
    b := o.MultiplyScalar(math.Sqrt(discriminant))
    return true, a.Sub(b)
  }
  return false, Vector{}
}

func (v Vector) Cross(o Vector) Vector {
  a := v.Y*o.Z - v.Z*o.Y
  b := v.Z*o.X - v.X*o.Z
  c := v.X*o.Y - v.Y*o.X
  return Vector{a, b, c}
}