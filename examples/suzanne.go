package main

import (
	"log"

	"github.com/kirillrdy/pt/pt"
)

func main() {
	scene := pt.Scene{}
	material := pt.DiffuseMaterial(pt.HexColor(0x334D5C))
	scene.Add(pt.NewSphere(pt.Vector{0.5, 1, 3}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{1.5, 1, 3}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewCube(pt.Vector{-5, -5, -2}, pt.Vector{5, 5, -1}, material))
	mesh, err := pt.LoadOBJ("examples/suzanne.obj", pt.SpecularMaterial(pt.HexColor(0xEFC94C), 2))
	if err != nil {
		log.Fatalln("LoadOBJ error:", err)
	}
	scene.Add(mesh)
	camera := pt.LookAt(pt.Vector{1, -0.45, 4}, pt.Vector{1, -0.6, 0.4}, pt.Vector{0, 1, 0}, 45)
	pt.RenderToWindow(pt.Render(&scene, &camera))
}
