package main

import "fmt"

func main() {
	slice := make([]int, 0, 100)
	oddSlice := fill(slice, 100)
	fmt.Println(oddSlice)

	done := make(chan struct{})
	sqSlice := squareConcurrently(oddSlice, done)
	fmt.Println(sqSlice)
}

func squareConcurrently(slice []int, done chan struct{}) []int {
	for i, val := range slice {
		i := i
		val := val
		go func() {
			slice[i] = val * val
			if i == len(slice)-1 {
				done <- struct{}{}
			}
		}()
	}
	<-done
	return slice
}

func fill(slice []int, count int) []int {
	for i := 0; i < count; i++ {
		slice = append(slice, i*2)
	}
	return slice
}
