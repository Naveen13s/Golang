package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Err   error
}

// func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {
func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()

	for job := range jobsChan {
		time.Sleep(time.Millisecond * 50)
		//fmt.Printf("image processed: %s\n", job)

		resultChan <- Result{
			Value: job,
			Err:   nil,
		}

	}
	fmt.Printf("Worker shutting down")

}

/*
	//time.Sleep(time.Millisecond * 50)

	//fmt.Printf("image processed: %s\n", url)

	resultChan <- Result{
		Value: url,
		Err:   nil,
	}
}

/*
// without concurrency normal - it took time in 100 ms to run 2 worker.
func main() {
	startTime := time.Now()
	result1 := worker("image_1.png")
	result2 := worker("image_2.png")
	worker(("image_3.png"))
	fmt.Println("result", result1)
	fmt.Println("result", result2)
	fmt.Printf("it took %s ms.", time.Since((startTime)))
}    */

/*
// it took 51.8174ms ms to run for 5 worker.
// Fan OUT:

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	wg.Add(5)
	go worker("image_1.png", &wg)
	go worker("image_2.png", &wg)
	go worker("image_3.png", &wg)
	go worker("image_4.png", &wg)
	go worker("image_5.png", &wg)
	//fmt.Println("result", result1)
	//fmt.Println("result", result2)
	wg.Wait()

	fmt.Printf("it took %s ms.", time.Since((startTime)))
}   */

/*

//Channel (FIFO- Queu) - buffered
//Fan in/out

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan Result, 5)

	startTime := time.Now()

	wg.Add(5)
	go worker("image_1.png", &wg, resultChan)
	go worker("image_2.png", &wg, resultChan)
	go worker("image_3.png", &wg, resultChan)
	go worker("image_4.png", &wg, resultChan)
	go worker("image_5.png", &wg, resultChan)

	wg.Wait()
	close(resultChan) // important  to close a channel, otherwise it throw fatal error: all goroutines are asleep - deadlock!

	for result := range resultChan {
		fmt.Printf("received: %v\n", result)

	}

	fmt.Printf("it took %v ms.", time.Since((startTime)))

}   */

// Launch and wait
//light weigh thread - > worker

// worker pool pattern
// it took  only 52.0594ms ms to  16 imageg process.
func main() {
	jobs := []string{
		"image_1.png",
		"image_2.png",
		"image_3.png",
		"image_4.png",
		"image_5.png",
		"image_6.png",
		"image_7.png",
		"image_9.png",
		"image_10.png",
		"image_11.png",
		"image_12.png",
		"image_13.png",
		"image_14.png",
		"image_15.png",
		"image_16.png",
	}
	var wg sync.WaitGroup
	totalWorkers := 5

	resultChan := make(chan Result, 50)

	jobsChan := make(chan string, len(jobs))

	startTime := time.Now()
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan) // now we will not send job first.
	}

	go func() {
		wg.Wait()
		close(resultChan) // important  to close a channel, otherwise it throw fatal error: all goroutines are asleep - deadlock!
	}()
	// send the jobs
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}
	close(jobsChan)

	for result := range resultChan {
		fmt.Printf(" Job completed: %v\n", result)

	}

	fmt.Printf("it took %v ms.", time.Since((startTime)))

}
