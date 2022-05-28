package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	count int
}

func (c *counter) add() {
	c.Lock()
	c.count++
	c.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	cnt := counter{count: 5}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				cnt.add()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(cnt.count)
}
