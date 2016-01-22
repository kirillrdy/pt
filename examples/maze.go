package main

import (
	"math/rand"

	"github.com/kirillrdy/pt/pt"
	"github.com/kirillrdy/render"
)

func uint8ToPixel(val uint8) int {
	fraction := float64(val) / 255.0
	return int(fraction * 65555)
}

func main() {

	width := 640
	height := 480

	render.Init()

	scene := pt.Scene{}
	floor := pt.GlossyMaterial(pt.HexColor(0x7E827A), 1.1, pt.Radians(30))
	material := pt.GlossyMaterial(pt.HexColor(0xE3CDA4), 1.1, pt.Radians(30))
	scene.Add(pt.NewCube(pt.Vector{-10000, -10000, -10000}, pt.Vector{10000, 10000, 0}, floor))
	n := 24
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			if rand.Float64() > 0.8 {
				min := pt.Vector{float64(x) - 0.5, float64(y) - 0.5, 0}
				max := pt.Vector{float64(x) + 0.5, float64(y) + 0.5, 1}
				cube := pt.NewCube(min, max, material)
				scene.Add(cube)
			}
		}
	}
	a := pt.NoAttenuation // QuadraticAttenuation(0.25)
	scene.Add(pt.NewSphere(pt.Vector{0, 0, 2.25}, 0.25, pt.LightMaterial(pt.Color{1, 1, 1}, 1, a)))
	camera := pt.LookAt(pt.Vector{1, 0, 30}, pt.Vector{0, 0, 0}, pt.Vector{0, 0, 1}, 35)
	//IterativeRender("out%03d.png", 1000, &scene, &camera, 2560, 1440, -1, 4, 4)

	cameraSamples := 200
	hitSamples := 200
	bounces := 10

	eventsChan := pt.Render(&scene, &camera, width, height, cameraSamples, hitSamples, bounces)

	for {
		event := <-eventsChan
		avg := event.Pixel
		render.SetPixel(event.X, event.Y, uint8ToPixel(avg.R), uint8ToPixel(avg.G), uint8ToPixel(avg.B))
	}
}

//TODO write a function that averages 2 images
//
//pixels := make([]pt.Color, width*height)
//for i := 1; i <= iterations; i++ {
//	fmt.Printf("\n[Iteration %d of %d]\n", i, iterations)
//	frame := pt.Render(&scene, &camera, width, height, cameraSamples, hitSamples, bounces)
//	for y := 0; y < height; y++ {
//		for x := 0; x < width; x++ {
//			index := y*width + x
//			c := pt.NewColor(frame.At(x, y))
//			pixels[index] = pixels[index].Add(c)
//			avg := pixels[index].DivScalar(float64(i))
//
//			renderer.SetDrawColor(avg.RGBA().R, avg.RGBA().G, avg.RGBA().B, 255)
//			renderer.DrawPoint(x, y)
//			renderer.Present()
//		}
//	}
//}
