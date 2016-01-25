package main

import (
	"github.com/kirillrdy/pt/pt"
	"github.com/ojrac/opensimplex-go"
)

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.Color{1, 1, 1}, 1.2, pt.Radians(20))
	noise := opensimplex.New()
	n := 80
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n*2; k++ {
				x := float64(i - n/2)
				y := float64(j - n/2)
				z := float64(k)
				m := 0.15
				w := noise.Eval3(x*m, y*m, z*m)
				w = (w + 0.8) / 1.6
				if w <= 0.2 {
					shape := pt.NewSphere(pt.Vector{x, y, z}, 0.333, material)
					scene.Add(shape)
				}
			}
		}
	}
	light := pt.NewSphere(pt.Vector{100, 0, 50}, 5, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation))
	scene.Add(light)
	camera := pt.LookAt(pt.Vector{0, 0, -20}, pt.Vector{}, pt.Vector{0, 1, 0}, 30)
	pt.RenderToWindow(pt.Render(&scene, &camera))
}
