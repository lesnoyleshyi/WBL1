package main

import "fmt"

func main() {
	numbers := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, val := range numbers {
			ch1 <- val
		}
		close(ch1)
	}()

	go func() {
		for val := range ch1 {
			ch2 <- val * val
		}
		close(ch2)
	}()

	for sqVal := range ch2 {
		fmt.Printf("%d ", sqVal)
	}
	fmt.Println()
}
