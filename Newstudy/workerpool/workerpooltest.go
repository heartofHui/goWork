package main

import (
	"fmt"
	"time"
)

//使用 goroutine 开启大小为3的线程池
//两个channel,1个为工作对列，

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, " processing job:", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 0; a < 9; a++ {
		<-results
	}
}
