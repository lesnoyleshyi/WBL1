package main

import (
	"fmt"
	"time"
)

func Sleep(d int64) {
	start := time.Now().UnixMilli()
	for time.Now().UnixMilli()-start < d {
	}
}

func main() {
	fmt.Println("Start: ", time.Now().Unix())

	Sleep(3000) //Sleep in milliseconds

	fmt.Println("End:   ", time.Now().Unix())
}