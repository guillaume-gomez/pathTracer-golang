package rendering

import (
  p_ "../primitives"
)


func PlaneScene() p_.World {
  world := p_.World{}
  plane := p_.NewPlane(p_.Vector{0, 0, 4}, p_.Vector{4,5,-1}, p_.Lambertian{p_.Vector{0.1, 0.3, 0.8}})
  world.AddAll(&plane)
  return world
}

func DiskScene() p_.World {
  world := p_.World{}
  disk := p_.NewDisk(p_.Vector{0, 0, 45}, p_.Vector{0,0,-1}, p_.Lambertian{p_.Vector{0.1, 0.3, 0.8}}, 100)
  world.AddAll(&disk)
  return world
}

func OriginalScene() p_.World {	
  world  := p_.World{}
  sphere := p_.NewSphere(0, 0, -1, 0.5, p_.Lambertian{p_.Vector{0.8, 0.3, 0.3}})
  floor  := p_.NewSphere(0, -100.5, -1, 100, p_.Lambertian{p_.Vector{0.8, 0.8, 0.0}})
  front  := p_.NewSphere(0, 0, 1, 0.2, p_.Lambertian{p_.Vector{0.8, 0.3, 0.3}})
  metal  := p_.NewSphere(1, 0, -1, 0.5, p_.Metal{p_.Vector{0.8, 0.6, 0.2}, 0.3})
  glass  := p_.NewSphere(-1, 0, -1, 0.5, p_.Dielectric{1.5})
  bubble := p_.NewSphere(-1, 0, -1, -0.45, p_.Dielectric{1.5})
  world.AddAll(&sphere, &floor, &front, &metal, &glass, &bubble)
  return world
}