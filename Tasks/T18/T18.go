package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count int64
}

func (c *counter) add(count int64) {
	atomic.AddInt64(&c.count, count)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	cnt := counter{count: 5}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				cnt.add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(cnt.count)
}
