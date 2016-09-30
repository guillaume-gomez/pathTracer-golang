package primitives

type Material interface {
  Bounce(input Ray, hit HitRecord) (bool, Ray)
  Color() Vector
}

func reflect(v Vector, n Vector) Vector {
  b := 2 * v.Dot(n)
  return v.Sub(n.MultiplyScalar(b))
}