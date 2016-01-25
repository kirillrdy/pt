package main

import (
	"github.com/kirillrdy/pt/pt"
)

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xEFC94C), 3, pt.Radians(30))
	whiteMat := pt.GlossyMaterial(pt.Color{1, 1, 1}, 3, pt.Radians(30))
	for x := 0; x < 10; x++ {
		for z := 0; z < 10; z++ {
			center := pt.Vector{float64(x) - 4.5, 0, float64(z) - 4.5}
			scene.Add(pt.NewSphere(center, 0.4, material))
		}
	}
	scene.Add(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, whiteMat))
	scene.Add(pt.NewSphere(pt.Vector{-1, 3, -1}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{0, 4, -8}, pt.Vector{0, 0, -2}, pt.Vector{0, 1, 0}, 45)
	pt.RenderToWindow(pt.Render(&scene, &camera))
}
