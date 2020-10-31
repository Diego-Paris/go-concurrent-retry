package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job alias
type Job = func() (string, error)

// Result contains the index of a task and its result
type Result struct {
	index  int
	result string
	retries int
}

func main() {
	fmt.Println("this is not control group!")
	start := time.Now()

	var tasks []Job

	for i := 0; i < 20; i++ {
		tasks = append(tasks, coding)
	}

	results := ConcurrentRetry(tasks, 3, 5)

	for r := range results {
		fmt.Println("Result received from thread:", r)
	}

	elapsed := time.Since(start)
	fmt.Println("Program finished in", elapsed)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Fake coding task
func coding() (string, error) {
	min := -5
	max := 5
	randVal := rand.Intn(max-min+1) + min

	time.Sleep(time.Millisecond)
	if randVal < 0 {
		return "Error!", errors.New("Error ocurred")
	}

	return "Success!", nil
}

func worker(id int, threads <-chan Job, results chan<- Result, retry int, wg *sync.WaitGroup) {
	//TODO

	//for each job in the threads channel
	for job := range threads {	
		
		var r Result

		for i := 1; i <= retry; i++ {
			
			res, err := job()

			// save the results even if an error occurred
			r = Result{
				index: id,
				result: res,
				retries: i,
			}

			// if no error ocurred, break out of loop
			if err == nil {
				break
			}
		}
		
		results <- r
		
	}

	wg.Done()
	
}

// ConcurrentRetry does stuff and things
func ConcurrentRetry(tasks []Job, concurrent, retry int) <-chan Result {

	threads := make(chan Job, len(tasks))
	results := make(chan Result, len(tasks))

	var wg sync.WaitGroup

	go func() {

		// Start sending jobs to the thread channel
		for _, task := range tasks {
			threads <- task
		}

		//fmt.Println("len", len(threads))
		//fmt.Println("cap", cap(threads))
		

		// Setup all the workers
		for i := 1; i <= concurrent; i++ {
			//fmt.Println("wow", i)
			wg.Add(1)
			go worker(i, threads, results, retry, &wg)
		}

		// loop manager
		for {
			if len(threads) == 0 {
				close(threads)
				//fmt.Println("broke out of loop")
				break
			}
		}

		wg.Wait()
		//close(threads)
		close(results)
	}()



	return results
}
