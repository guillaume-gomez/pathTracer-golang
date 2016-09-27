package primatives

type Material interphase {
	Bounce(input Ray, hit Hit) (bool, Ray)
	Color() Vector
}