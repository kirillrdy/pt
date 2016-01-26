package pt

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func Render(scene *Scene, camera *Camera) chan pixelRenderResult {
	pixelJobs := pixelJobsAtRandom(RenderConfig.Width, RenderConfig.Height)
	scene.Compile()
	absCameraSamples := int(math.Abs(float64(RenderConfig.CameraSamples)))
	fmt.Printf("%d x %d pixels, %d x %d = %d samples, %d bounces \n",
		RenderConfig.Width, RenderConfig.Height, absCameraSamples, RenderConfig.HitSamples, absCameraSamples*RenderConfig.HitSamples, RenderConfig.Bounces)
	scene.rays = 0
	results := make(chan pixelRenderResult)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for pixelJob := range pixelJobs {
				renderEvent := pixelRender(scene, camera, pixelJob.x, pixelJob.y, absCameraSamples)
				results <- renderEvent
			}
			//TODO find a way of closing results when we ran out of jobs
			//close(results)
		}()
	}
	return results
}

func pixelRender(scene *Scene, camera *Camera, x, y int, absCameraSamples int) pixelRenderResult {

	c := Color{}
	if RenderConfig.CameraSamples <= 0 {
		// random subsampling
		for i := 0; i < absCameraSamples; i++ {
			fu := rand.Float64()
			fv := rand.Float64()
			ray := camera.CastRay(x, y, fu, fv)
			c = c.Add(scene.Sample(ray, true, RenderConfig.HitSamples, RenderConfig.Bounces))
		}
		c = c.DivScalar(float64(absCameraSamples))
	} else {
		// stratified subsampling
		n := int(math.Sqrt(float64(RenderConfig.CameraSamples)))
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				fu := (float64(u) + 0.5) / float64(n)
				fv := (float64(v) + 0.5) / float64(n)
				ray := camera.CastRay(x, y, fu, fv)
				c = c.Add(scene.Sample(ray, true, RenderConfig.HitSamples, RenderConfig.Bounces))
			}
		}
		c = c.DivScalar(float64(n * n))
	}
	c = c.Pow(1 / 2.2)
	return pixelRenderResult{X: x, Y: y, Pixel: c}
}

//TODO restore this, using channels
// func IterativeRender(pathTemplate string, iterations int, scene *Scene, camera *Camera, w, h, cameraSamples, hitSamples, bounces int) error {
// 	scene.Compile()
// 	pixels := make([]Color, w*h)
// 	result := image.NewRGBA(image.Rect(0, 0, w, h))
// 	for i := 1; i <= iterations; i++ {
// 		fmt.Printf("\n[Iteration %d of %d]\n", i, iterations)
// 		frame := Render(scene, camera, w, h, cameraSamples, hitSamples, bounces)
// 		for y := 0; y < h; y++ {
// 			for x := 0; x < w; x++ {
// 				index := y*w + x
// 				c := NewColor(frame.At(x, y))
// 				pixels[index] = pixels[index].Add(c)
// 				avg := pixels[index].DivScalar(float64(i))
// 				result.SetRGBA(x, y, avg.RGBA())
// 			}
// 		}
// 		path := pathTemplate
// 		if strings.Contains(path, "%") {
// 			path = fmt.Sprintf(pathTemplate, i)
// 		}
// 		if err := SavePNG(path, result); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
