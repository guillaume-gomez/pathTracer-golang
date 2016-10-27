package main

import (
  "flag"
  "fmt"
  "math"
  "math/rand"
  "strconv"
  "os"
  "time"
  "strings"
  s_ "./rendering"
  p_ "./primitives"
)

const (
  c  = 255.99
  extension = ".ppm"
  progressBarWidth = 80
)

var config struct {
  nx, ny, ns    int
  aperture, fov float64
  filename      string
  lookFrom      p_.Vector
  debug         bool
}

var (
  white = p_.Vector{1.0, 1.0, 1.0}
  blue  = p_.Vector{0.5, 0.7, 1.0}
)

func check(err error, msg string) {
  if err != nil {
    fmt.Fprintf(os.Stderr, msg, err)
    os.Exit(1)
  }
}

func colorize(r p_.Ray, world p_.Hitable, depth int) p_.Vector {
  hit, record := world.Hit(r, 0.001, math.MaxFloat64)
  if hit {
    if depth < 50 {
      bounced, bouncedRay := record.Bounce(r, record)
      if bounced {
        newColor := colorize(bouncedRay, world, depth+1)
        return record.Material.Color().Mul(newColor)
      }
    }
    return p_.Vector{}
  }

  return gradient(r)
}

func colorizeSimplified(r p_.Ray, world p_.Hitable, depth int) p_.Vector {
  hit, _ := world.Hit(r, 0.001, math.MaxFloat64)
  if hit {
    return p_.Vector{1,0,0} //red
  }

  return gradient(r)
}


func gradient(r p_.Ray) p_.Vector {
  // make unit vector so y is between -1.0 and 1.0
  v := r.Direction().Normalize()
  // scale t to be between 0.0 and 1.0
  t := 0.5 * (v.Y + 1.0)

  // linear blend: blended_value = (1 - t) * white + t * blue
  return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func createFile(filename string) *os.File {
  f, err := os.Create(filename)
  check(err, "Error opening file: %v\n")

  // http://netpbm.sourceforge.net/doc/ppm.html
  _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", config.nx, config.ny)
  check(err, "Error writting to file: %v\n")
  return f
}

func writeFile(f *os.File, rgb p_.Vector) {
  // get inteconfig.nsity of colors
  ir := int(c * math.Sqrt(rgb.X))
  ig := int(c * math.Sqrt(rgb.Y))
  ib := int(c * math.Sqrt(rgb.Z))

  _, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
  check(err, "Error writing to file: %v\n")
}

// samples rays for anti-aliasing
func sample(world *p_.World, camera *p_.Camera, i, j int) p_.Vector {
  rgb := p_.Vector{}

  for s := 0; s < config.ns; s++ {
    u := (float64(i) + rand.Float64()) / float64(config.nx)
    v := (float64(j) + rand.Float64()) / float64(config.ny)

    r := camera.RayAt(u, v)
    color := p_.Vector{}
    if config.debug {
      color = colorizeSimplified(r, world, 0)
    } else {
      color = colorize(r, world, 0)
    }
    rgb = rgb.Add(color)
  }

  // average
  return rgb.DivideScalar(float64(config.ns))
}

func render(world *p_.World, camera *p_.Camera, file string) {
  ticker := time.NewTicker(time.Millisecond * 100)

  go func() {
    for {
      <-ticker.C
      fmt.Print(".")
    }
  }()

  f := createFile(file)
  defer f.Close()

  start := time.Now()

  for j := config.ny - 1; j >= 0; j-- {
    for i := 0; i < config.nx; i++ {
      rgb := sample(world, camera, i, j)
      writeFile(f, rgb)
    }
  }

  ticker.Stop()
  fmt.Printf("\nDone.\nElapsed: %v\n", time.Since(start))
}

func slowlyMoveBack(world p_.World, camera p_.Camera, nbImage int, step float64 ) {
  imageIndex := 0
  for imageIndex < nbImage {
    file := config.filename + "_" + strconv.Itoa(imageIndex) + "." + extension
    fmt.Printf("Begin computing : %s\n", file)
    render(&world, &camera, file)
    fmt.Print("--End--\n")
    camera.MoveTo(p_.Vector{0.0, 0.0, step * float64(imageIndex + 1)})
    imageIndex += 1
  }
}

func initCommandLineParams() {
  flag.Float64Var(&config.lookFrom.X, "x", 0, "look from X")
  flag.Float64Var(&config.lookFrom.Y, "y", 0, "look from Y")
  flag.Float64Var(&config.lookFrom.Z, "z", 0, "look from Z")

  flag.Float64Var(&config.fov, "fov", 90.0, "vertical field of view (degrees)")
  flag.IntVar(&config.nx, "width", 400, "width of image")
  flag.IntVar(&config.ny, "height", 200, "height of image")
  flag.IntVar(&config.ns, "samples", 100, "number of samples for anti-aliasing")
  flag.StringVar(&config.filename, "out", "out", "output filename")
  flag.BoolVar(&config.debug, "debug", false, "debug true")
  flag.Parse()

}

func main() {
  initCommandLineParams()
  aperture := 2.0
  lookAt := p_.Vector{0,0,-1}
  focusDist := config.lookFrom.Sub(lookAt).Length()


  camera := p_.NewCamera(config.lookFrom, lookAt, p_.Vector{0,1,0}, config.fov, float64(config.nx)/float64(config.ny), aperture, focusDist)
  world := s_.DiskScene()

  fmt.Printf("\nRendering %d x %d pixel scene with %d objects:", config.nx, config.ny, 6)
  fmt.Printf("\n[%d samples/pixel, %.2f° fov, %.2f aperture]\n", config.ns, config.fov, aperture)

  render(&world, &camera, config.filename + extension)
  //slowlyMoveBack(world, camera, filename, 10, 1.0)
}


func outputProgress(ch chan int, rows int) {
  fmt.Println()
  for i := 1; i <= rows; i++ {
    <-ch
    pct := 100 * float64(i) / float64(rows)
    filled := (progressBarWidth * i) / rows
    bar := strings.Repeat("=", filled) + strings.Repeat("-", progressBarWidth-filled)
    fmt.Printf("\r[%s] %.2f%%", bar, pct)
  }
  fmt.Println()
}