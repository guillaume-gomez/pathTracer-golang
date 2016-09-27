package primatives

type Lambertian struct {
	C Vector
}

func (l Lambertian) Bounce(input Ray, hit HitRecord) (bool, Ray) {
	direction := hit.Normal.Add(VectorInUnitSphere())
	return true, Ray{hit.P, direction}
}

func (l Lambertian) Color() Vector {
	return l.C
}