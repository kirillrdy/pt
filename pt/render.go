package pt

import (
	"fmt"
	"github.com/kirillrdy/pt/xlib"
	"log"
	"math"
	"math/rand"
	"runtime"
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

	xlib.CreateWindow(RenderConfig.Width, RenderConfig.Height)

	for event := range eventsChan {
		xlib.SetPixel(event.X, event.Y, int(event.Pixel.R*65535), int(event.Pixel.G*65535), int(event.Pixel.B*65535))
	}
}

func pixelJobsAtRandom(width int, height int) chan pixelJob {
	start := time.Now()
	jobs := make([]pixelJob, 0, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			jobs = append(jobs, pixelJob{x: x, y: y})
		}
	}

	for count := 0; count < len(jobs); count++ {
		a := rand.Intn(len(jobs))
		b := rand.Intn(len(jobs))
		jobs[a], jobs[b] = jobs[b], jobs[a]
	}

	jobChannel := make(chan pixelJob)
	log.Printf("Randomizing pixels took %v\n ", time.Since(start))
	go func() {
		for _, job := range jobs {
			jobChannel <- job
		}
		close(jobChannel)
	}()

	return jobChannel
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

func Render(scene *Scene, camera *Camera) chan ResultEvent {
	pixelJobs := pixelJobsAtRandom(RenderConfig.Width, RenderConfig.Height)
	scene.Compile()
	absCameraSamples := int(math.Abs(float64(RenderConfig.CameraSamples)))
	fmt.Printf("%d x %d pixels, %d x %d = %d samples, %d bounces \n",
		RenderConfig.Width, RenderConfig.Height, absCameraSamples, RenderConfig.HitSamples, absCameraSamples*RenderConfig.HitSamples, RenderConfig.Bounces)
	scene.rays = 0
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	results := make(chan ResultEvent)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			//TODO do we still need this
			//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for pixelJob := range pixelJobs {
				renderEvent := pixelRender(scene, camera, pixelJob.x, pixelJob.y, absCameraSamples)
				results <- renderEvent
			}
			//close(results)
		}()
	}
	return results
}

func pixelRender(scene *Scene, camera *Camera, x, y int, absCameraSamples int) ResultEvent {

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
