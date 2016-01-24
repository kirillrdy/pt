package pt

import (
	"fmt"
	"github.com/kirillrdy/pt/xlib"
	"math"
	"math/rand"
	"time"
)

func showProgress(start time.Time, rays uint64, i, h int) {
	pct := int(100 * float64(i) / float64(h))
	elapsed := time.Since(start)
	rps := float64(rays) / elapsed.Seconds()
	fmt.Printf("\r%4d / %d (%3d%%) [", i, h, pct)
	for p := 0; p < 100; p += 3 {
		if pct > p {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %s %s ", DurationString(elapsed), NumberString(rps))
}

type ResultEvent struct {
	X     int
	Y     int
	Pixel Color
}

type pixelJob struct {
	x int
	y int
}

func RenderToWindow(eventsChan chan ResultEvent) {

	for event := range eventsChan {
		xlib.SetPixel(event.X, event.Y, int(event.Pixel.R*65535), int(event.Pixel.G*65535), int(event.Pixel.B*65535))
	}
}

func pixelJobs(width int, height int) chan pixelJob {
	jobChannel := make(chan pixelJob)
	go func() {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				jobChannel <- pixelJob{x: x, y: y}
			}
		}
		close(jobChannel)
	}()
	return jobChannel
}

func Render(scene *Scene, camera *Camera, w, h, cameraSamples, hitSamples, bounces int) chan ResultEvent {

	pixelJobs := pixelJobs(w, h)

	scene.Compile()
	absCameraSamples := int(math.Abs(float64(cameraSamples)))
	fmt.Printf("%d x %d pixels, %d x %d = %d samples, %d bounces \n",
		w, h, absCameraSamples, hitSamples, absCameraSamples*hitSamples, bounces)
	scene.rays = 0
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	results := make(chan ResultEvent)
	go func() {
		for pixelJob := range pixelJobs {
			x := pixelJob.x
			y := pixelJob.y
			renderEvent := pixelRender(w, h, scene, camera, x, y, absCameraSamples, cameraSamples, hitSamples, bounces, rnd)
			results <- renderEvent
		}

		close(results)
	}()
	return results
}

func pixelRender(w, h int, scene *Scene, camera *Camera, x, y int, absCameraSamples, cameraSamples, hitSamples, bounces int, rnd *rand.Rand) ResultEvent {

	c := Color{}
	if cameraSamples <= 0 {
		// random subsampling
		for i := 0; i < absCameraSamples; i++ {
			fu := rnd.Float64()
			fv := rnd.Float64()
			ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
			c = c.Add(scene.Sample(ray, true, hitSamples, bounces, rnd))
		}
		c = c.DivScalar(float64(absCameraSamples))
	} else {
		// stratified subsampling
		n := int(math.Sqrt(float64(cameraSamples)))
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				fu := (float64(u) + 0.5) / float64(n)
				fv := (float64(v) + 0.5) / float64(n)
				ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
				c = c.Add(scene.Sample(ray, true, hitSamples, bounces, rnd))
			}
		}
		c = c.DivScalar(float64(n * n))
	}
	c = c.Pow(1 / 2.2)
	return ResultEvent{X: x, Y: y, Pixel: c}
}

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
