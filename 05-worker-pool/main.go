package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	JobId int
}

func Worker(workerNo int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Job %d picked up by Worker %d\n", job.JobId, workerNo)
		timeInMillis := rand.Float32() * 1000
		time.Sleep(time.Duration(timeInMillis) * time.Millisecond)
	}
}

func main() {

	jobs := make(chan Job)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, jobs, &wg)
	}

	for i := 1; i <= 200; i++ {
		jobs <- Job{
			JobId: i,
		}
	}

	close(jobs)

	wg.Wait()

}
