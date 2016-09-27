package util

import (
  p_ "./../primatives"
  "math"
  "math/rand"
  "fmt"
  "os"
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

  sphere = p_.Sphere{p_.Vector{0, 0, -1}, 0.5}
  floor  = p_.Sphere{p_.Vector{0, -100.5, -1}, 100}
  world = p_.World{[]p_.Hitable{&sphere, &floor}}

)

func colorize(r p_.Ray, world p_.Hitable, depth int) p_.Vector {
  hit, record := world.Hit(r, 0.001, math.MaxFloat64)
  if hit {
    if depth < 50 {
      fmt.Printf("toto")
      bounced, bouncedRay := record.Bounce(r, record)
      fmt.Printf("fdjkjf\n")
      if bounced {
        newColor := colorize(bouncedRay, world, depth+1)
        return record.Material.Color().Mul(newColor)
      }
    }
    return p_.Vector{}
  }

  return gradient(r)
}


func gradient(r p_.Ray) p_.Vector {
  // make unit vector so y is between -1.0 and 1.0
  v := r.Direction.Normalize()
  // scale t to be between 0.0 and 1.0
  t := 0.5 * (v.Y + 1.0)

  // linear blend: blended_value = (1 - t) * white + t * blue
  return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func GradientSphere(x int, nx int, y int, ny int,f *os.File) error {
  rgb := p_.Vector{}
  // sample rays for anti-aliasing
    for s := 0; s < ns; s++ {
      u := (float64(x) + rand.Float64()) / float64(nx)
      v := (float64(y) + rand.Float64()) / float64(ny)

      r := camera.RayAt(u, v)
      color := colorize(r, &world, 0)
      rgb = rgb.Add(color)
    }

  // average
  rgb = rgb.DivideScalar(float64(ns))

  // get intensity of colors
  ir := int(color * rgb.X)
  ig := int(color * rgb.Y)
  ib := int(color * rgb.Z)

  _, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
  return err
}
