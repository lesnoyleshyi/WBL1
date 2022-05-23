package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("provide lifetime in seconds as the only argument")
	}
	lifetime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("can't parse lifetime: %s", err)
	}
	mainCh := make(chan string)
	done := make(chan struct{})
	go cancelTimeout(done, lifetime)
	go writeToCh(done, mainCh)
	go readFromCh(done, mainCh)

	<-done
	close(mainCh)
}

func writeToCh(done <-chan struct{}, ch chan string) {
	for {
		timestamp := fmt.Sprintf("it's %s", time.Now().Format(time.StampMicro))
		select {
		case <-done:
			return
		case ch <- timestamp:
			time.Sleep(time.Millisecond * 250)
		}
	}
}

func readFromCh(done <-chan struct{}, ch chan string) {
	for msg := range ch {
		select {
		case <-done:
			return
		default:
			fmt.Printf("new message: %s\n", msg)
		}
	}
}

func cancelTimeout(done chan struct{}, timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	log.Println("Timeout exceeded. Program canceled")
	done <- struct{}{}
}
