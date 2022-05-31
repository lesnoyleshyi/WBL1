package main

import (
	"fmt"
	"time"
)

func Sleep(d int64) {
	start := time.Now()
	end := start.Add(time.Millisecond*time.Duration(d))
	for time.Now().UnixNano() < end.UnixNano() { }
}

func main() {
	start := time.Now()
	fmt.Println("Start: ", start)

	Sleep(3000) //Sleep in milliseconds

	end := time.Now()
	fmt.Println("End:   ", end)
	fmt.Printf("Time delta is %d millisec\n", end.Sub(start))
}