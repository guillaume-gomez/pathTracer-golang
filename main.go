package main

import (
  "fmt"
  "os"
  p_ "./primatives"
  utils "./util"
)

const (
  nx = 400 // size of x
  ny = 200 // size of y
  ns = 100 // number of samples for aa
  color  = 255.99
)

var (
  white = p_.Vector{ 1.0, 1.0, 1.0 }
  blue  = p_.Vector{ 0.5, 0.7, 1.0 }

  camera = p_.NewCamera()

  sphere = p_.Sphere{ p_.Vector{ 0, 0, -1 }, 0.5 }
  floor  = p_.Sphere{ p_.Vector{ 0, -100.5, -1 }, 100 }
)

func check(e error, s string) {
  if e != nil {
    fmt.Fprintf(os.Stderr, s, e)
    os.Exit(1)
  }
}

func main() {
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
      err = utils.GradientSphere(i, nx, j, ny, f)
      check(err, "Error writing to file: %v\n")
    }
  }
}