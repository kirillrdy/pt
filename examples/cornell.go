package main

import "github.com/kirillrdy/pt/pt"
import "github.com/kirillrdy/pt/xlib"

func main() {
	white := pt.DiffuseMaterial(pt.Color{0.740, 0.742, 0.734})
	red := pt.DiffuseMaterial(pt.Color{0.366, 0.037, 0.042})
	green := pt.DiffuseMaterial(pt.Color{0.163, 0.409, 0.083})
	light := pt.LightMaterial(pt.Color{0.780, 0.780, 0.776}, 10, pt.QuadraticAttenuation(0.1))
	scene := pt.Scene{}
	n := 10.0
	scene.Add(pt.NewCube(pt.Vector{-n, -11, -n}, pt.Vector{n, -10, n}, white))
	scene.Add(pt.NewCube(pt.Vector{-n, 10, -n}, pt.Vector{n, 11, n}, white))
	scene.Add(pt.NewCube(pt.Vector{-n, -n, 10}, pt.Vector{n, n, 11}, white))
	scene.Add(pt.NewCube(pt.Vector{-11, -n, -n}, pt.Vector{-10, n, n}, red))
	scene.Add(pt.NewCube(pt.Vector{10, -n, -n}, pt.Vector{11, n, n}, green))
	scene.Add(pt.NewSphere(pt.Vector{3, -7, -3}, 3, white))
	cube := pt.NewCube(pt.Vector{-3, -4, -3}, pt.Vector{3, 4, 3}, white)
	transform := pt.Rotate(pt.Vector{0, 1, 0}, pt.Radians(30)).Translate(pt.Vector{-3, -6, 4})
	scene.Add(pt.NewTransformedShape(cube, transform))
	scene.Add(pt.NewCube(pt.Vector{-2, 9.8, -2}, pt.Vector{2, 10, 2}, light))
	camera := pt.LookAt(pt.Vector{0, 0, -20}, pt.Vector{0, 0, 1}, pt.Vector{0, 1, 0}, 65)

	width := 512
	height := 512
	xlib.CreateWindow(width, height)
	pt.RenderToWindow(pt.Render(&scene, &camera, width, height))
}
