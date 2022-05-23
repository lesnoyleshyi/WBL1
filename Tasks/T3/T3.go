package main

import (
	"fmt"
	"sync"
)

func main() {
	var workersCnt int = 10

	//generate input
	src := make([]int, 100)
	for i := 0; i < 100; i++ {
		src[i] = i * 2
	}
	nums := gen(src...)

	//square numbers concurrently by defined number of workers
	squaresSlice := make([]<-chan int, 0, workersCnt)
	for i := 0; i < workersCnt; i++ {
		sq := square(nums)
		squaresSlice = append(squaresSlice, sq)
	}

	squares := merge(squaresSlice...)

	fmt.Println(sum(squares))

}

func gen(input ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range input {
			out <- num
		}
		close(out)
	}()
	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range input {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	drain := func(ch <-chan int) {
		for val := range ch {
			out <- val
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go drain(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func sum(input <-chan int) int {
	var res int
	for num := range input {
		res += num
	}

	return res
}