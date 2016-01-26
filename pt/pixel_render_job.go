package pt

import (
	"log"
	"math/rand"
	"time"
)

type pixelRenderJob struct {
	x int
	y int
}

func pixelJobsAtRandom(width int, height int) chan pixelRenderJob {
	start := time.Now()
	jobs := make([]pixelRenderJob, 0, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			jobs = append(jobs, pixelRenderJob{x: x, y: y})
		}
	}

	for count := 0; count < len(jobs); count++ {
		a := rand.Intn(len(jobs))
		b := rand.Intn(len(jobs))
		jobs[a], jobs[b] = jobs[b], jobs[a]
	}

	jobChannel := make(chan pixelRenderJob)
	log.Printf("Randomizing pixels took %v\n ", time.Since(start))
	go func() {
		for _, job := range jobs {
			jobChannel <- job
		}
		close(jobChannel)
	}()

	return jobChannel
}

func pixelJobs(width int, height int) chan pixelRenderJob {
	jobChannel := make(chan pixelRenderJob)
	go func() {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				jobChannel <- pixelRenderJob{x: x, y: y}
			}
		}
		close(jobChannel)
	}()
	return jobChannel
}
