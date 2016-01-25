package main

import (
	"log"

	"github.com/kirillrdy/pt/pt"
)

func main() {
	scene := pt.Scene{}
	scene.Add(pt.NewSphere(pt.Vector{-2, 5, -3}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{5, 5, -3}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, pt.SpecularMaterial(pt.HexColor(0xFCFAE1), 2)))
	mesh, err := pt.LoadOBJ("examples/teapot.obj", pt.SpecularMaterial(pt.HexColor(0xB9121B), 2))
	if err != nil {
		log.Fatalln("LoadOBJ error:", err)
	}
	scene.Add(mesh)
	camera := pt.LookAt(pt.Vector{2, 5, -6}, pt.Vector{0.5, 1, 0}, pt.Vector{0, 1, 0}, 45)
	pt.RenderToWindow(pt.Render(&scene, &camera))
}
