package main

import (
	"log"
	"math/rand"

	"github.com/kirillrdy/pt/pt"
	"github.com/kirillrdy/pt/xlib"
)

func createMesh(material pt.Material) pt.Shape {
	mesh, err := pt.LoadBinarySTL("examples/cube.stl", material)
	if err != nil {
		log.Fatalln("LoadBinarySTL error:", err)
	}
	mesh.FitInside(pt.Box{pt.Vector{0, 0, 0}, pt.Vector{1, 1, 1}}, pt.Vector{0.5, 0.5, 0.5})
	return mesh
}

func main() {
	scene := pt.Scene{}
	meshes := []pt.Shape{
		createMesh(pt.GlossyMaterial(pt.HexColor(0x3B596A), 1.5, pt.Radians(20))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0x427676), 1.5, pt.Radians(20))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0x3F9A82), 1.5, pt.Radians(20))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xA1CD73), 1.5, pt.Radians(20))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xECDB60), 1.5, pt.Radians(20))),
	}
	for x := -8; x <= 8; x++ {
		for z := -12; z <= 12; z++ {
			fx := float64(x)
			fy := rand.Float64() * 2
			fz := float64(z)
			scene.Add(pt.NewTransformedShape(meshes[rand.Intn(len(meshes))], pt.Translate(pt.Vector{fx, fy, fz})))
			scene.Add(pt.NewTransformedShape(meshes[rand.Intn(len(meshes))], pt.Translate(pt.Vector{fx, fy - 1, fz})))
		}
	}
	scene.Add(pt.NewSphere(pt.Vector{8, 10, 0}, 3, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-10, 10, 0}, pt.Vector{-2, 0, 0}, pt.Vector{0, 1, 0}, 45)
	width := 2560 / 2
	height := 1440 / 2
	xlib.CreateWindow(width, height)
	pt.RenderConfig.CameraSamples = 40
	pt.RenderToWindow(pt.Render(&scene, &camera, width, height))
}
