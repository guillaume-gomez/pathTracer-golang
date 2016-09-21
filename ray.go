package main


type Ray struct {
  Origin, Direction Vector
}


func (r Ray) Point(t float64) Vector {
	// p(t)  =  A + B * t
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}