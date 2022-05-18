package main

import "fmt"

func main() {
	src := make([]int, 100, 100)
	for i, _ := range src {
		src[i] = 2 * i
	}

	numbers := gen(src...)
	squares := square(numbers)
	for val := range squares {
		fmt.Println(val)
	}
}

func gen(values ...int) <-chan int {
	nums := make(chan int)
	go func() {
		for _, val := range values {
			nums <- val
		}
		close(nums)
	}()
	return nums
}

func square(nums <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		for num := range nums {
			squares <- numgit * num
		}
		close(squares)
	}()
	return squares
}
