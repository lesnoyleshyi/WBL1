package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	mainCh := make(chan string)
	//Send data to "main" channel permanently
	go func() {
		for {
			mainCh <- time.Now().Format(time.StampMilli)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	//get count of workers as the only parameter
	if len(os.Args) < 2 {
		log.Fatalln("Provide workers count as an argument, please")
	}
	workerCnt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("unknown count of workers: %s", err)
	}

	for i := 0; i < workerCnt; i++ {
		go work(i, mainCh)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
	close(mainCh)
}

func work(id int, ch chan string) {
	for smth := range ch {
		fmt.Printf("worker N%v read from channel at %s\n", id, smth)
	}
}
