package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, " processing job:", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	var worker_queues = 4
	//var worker_threads = 4
	var jobs_queues [] chan int
	jobs_queues = make([](chan int), worker_queues, 100)
	results := make(chan int, 100)

	fmt.Println(len(jobs_queues))
	fmt.Println(len(jobs_queues[0]))
	for w := 0; w < worker_queues; w++ {
		fmt.Println("go")
		go worker(w, jobs_queues[w], results)
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("push job:%d\n", j%worker_queues)
		jobs_queues[j%worker_queues] <- j
	}

	for i := 0; i < worker_queues; i++ {
		close(jobs_queues[i])
	}
}
