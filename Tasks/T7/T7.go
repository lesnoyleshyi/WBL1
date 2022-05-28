package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TrSafeMap struct {
	sync.RWMutex
	m map[string]int
}

func NewTreadSafeMap() TrSafeMap {
	m := make(map[string]int)
	return TrSafeMap{
		m: m,
	}
}

func (m *TrSafeMap) Add(key string, val int) {
	m.Lock()
	m.m[key] = val
	m.Unlock()
}

func (m TrSafeMap) Len() int {
	return len(m.m)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)

	m1 := NewTreadSafeMap()
	m2 := sync.Map{}
	//bad idea
	counter := 0

	for i := 0; i < 100; i++ {
		go saveToMyMap(&m1, &wg)
		go saveToSyncMap(&m2, &wg, &counter)
	}
	wg.Wait()
	fmt.Printf("All data is written to maps. len(m1)=%d; len(m2)=%d\n", m1.Len(), counter)
	//fmt.Println(m1.m)
	//fmt.Println(m2)
}

func randValsGen() (string, int) {
	names := []string{"James", "Mary", "Robert", "Patricia", "John", "Jennifer",
		"Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan",
		"Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen", "Max"}

	rand.Seed(time.Now().UnixNano())
	height := 150 + rand.Intn(50)
	name := names[rand.Intn(len(names))]

	return name, height
}

func saveToMyMap(m *TrSafeMap, wg *sync.WaitGroup) {
	name, height := randValsGen()
	m.Add(name, height)
	wg.Done()
}

func saveToSyncMap(m *sync.Map, wg *sync.WaitGroup, counter *int) {
	name, height := randValsGen()
	m.Store(name, height)
	*counter++
	wg.Done()
}
