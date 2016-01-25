package main

import (
	"log"

	"github.com/kirillrdy/pt/pt"
)

func createMesh(material pt.Material) pt.Shape {
	mesh, err := pt.LoadBinarySTL("examples/cylinder.stl", material)
	if err != nil {
		log.Fatalln("LoadBinarySTL error:", err)
	}
	mesh.FitInside(pt.Box{pt.Vector{-0.1, -0.1, 0}, pt.Vector{1.1, 1.1, 100}}, pt.Vector{0.5, 0.5, 0})
	mesh.SmoothNormalsThreshold(pt.Radians(10))
	return mesh
}

func main() {
	scene := pt.Scene{}
	meshes := []pt.Shape{
		createMesh(pt.GlossyMaterial(pt.HexColor(0x730046), 1.6, pt.Radians(45))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xBFBB11), 1.6, pt.Radians(45))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xFFC200), 1.6, pt.Radians(45))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xE88801), 1.6, pt.Radians(45))),
		createMesh(pt.GlossyMaterial(pt.HexColor(0xC93C00), 1.6, pt.Radians(45))),
	}
	for x := -6; x <= 3; x++ {
		mesh := meshes[(x+6)%len(meshes)]
		for y := -5; y <= 4; y++ {
			fx := float64(x) / 2
			fy := float64(y)
			fz := float64(x) / 2
			scene.Add(pt.NewTransformedShape(mesh, pt.Translate(pt.Vector{fx, fy, fz})))
		}
	}
	scene.Add(pt.NewSphere(pt.Vector{1, 0, 10}, 3, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-5, 0, 5}, pt.Vector{1, 0, 0}, pt.Vector{0, 0, 1}, 45)
	pt.RenderToWindow(pt.Render(&scene, &camera))
}
