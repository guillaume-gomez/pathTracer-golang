package main

import (
  "fmt"
  "os"
  "math"
)

const (
  color = 255.99
)

var (
  white = Vector{1.0, 1.0, 1.0}
  blue  = Vector{0.5, 0.7, 1.0}
  sphere = Sphere{Vector{0, 0, -1}, 0.5}
  lowerLeft = Vector{-2.0, -1.0, -1.0}
  horizontal = Vector{4.0, 0.0, 0.0}
  vertical = Vector{0.0, 2.0, 0.0}
  origin = Vector{0.0, 0.0, 0.0}
)

func check(e error, s string) {
  if e != nil {
    fmt.Fprintf(os.Stderr, s, e)
    os.Exit(1)
  }
}

func colorize(r *Ray, sphere Sphere) Vector {
  hit, record := sphere.Hit(r, 0.0, math.MaxFloat64)

  if hit {
    return record.Normal.AddScalar(1.0).MultiplyScalar(0.5)
  }

  // make unit vector so y is between -1.0 and 1.0
  unitDirection := r.Direction.Normalize()

  return gradient(&unitDirection)
}

func gradient(v *Vector) Vector {
  // scale t to be between 0.0 and 1.0
  t := 0.5 * (v.Y + 1.0)

  // linear blend: blended_value = (1 - t) * white + t * blue
  return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func gradientSphere(x int, nx int, y int, ny int,f *os.File) error {
  u := float64(x) / float64(nx)
  v := float64(y) / float64(ny)

  position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

  direction := lowerLeft.Add(position)
  r := Ray{origin, direction}
  rgb := colorize(&r, sphere)

  // get intensity of colors
  ir := int(color * rgb.X)
  ig := int(color * rgb.Y)
  ib := int(color * rgb.Z)

  _, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
  return err
}

func main() {
  // size of image x and y
  nx := 400
  ny := 200
  f, err := os.Create("out.ppm")

  defer f.Close()

  check(err, "Error opening file: %v\n")

  // http://netpbm.sourceforge.net/doc/ppm.html
  _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

  check(err, "Error writting to file: %v\n")
  // writes each pixel with r/g/b values
  // from top left to bottom right
  for j := ny - 1; j >= 0; j-- {
    for i := 0; i < nx; i++ {
      err = gradientSphere(i, nx, j, ny, f)
      check(err, "Error writing to file: %v\n")
    }
  }
}