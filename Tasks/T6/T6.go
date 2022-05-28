package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Горутину невозможно остановить принудительно извне (из другой горутины).
//Горутина должна завершаться самостоятельно.
//
//Отдельный случай - "аварийное" завершение всех горутин при завершении main-горутины.
//Также горутина завершается, когда выполнит все инструкции функции или выполнит вызов return явно.
//
//Общий паттерн для корректного (не аварийного) завершения - заставить горутину "слушать" канал.
//Как правило, этот канал называется done, но название ни на что не влияет.
//Канал может быть передан в функцию явно, либо быть в области видимости горутины.
//Чтение из канала имеет следующие особенности:
//	- выполнение горутины приостанавливается, если канал пуст. Горутина будет ждать, пока в канал кто-нибудь напишет.
//	- чтение из пустого канала не блокирует горутину, если чтение выполняется в конструкции select.
//	- чтение из закрытого канала не блокирует горутину.
//		(Если до этого канал был пуст, и горутина "висела", после закрытия канала она сразу отвиснет)
//На этих особенностях построены все методы принудительной остановки горутин.
//Главная идея всех методов - отправка в канал чего-угодно горутиной-убийцей с одной стороны
//								и чтение из этого канала горутиной-жертвой.
//Завершение горутины с помощью конекста также основано на чтении горутиной-жертвой данных из канала контекста.
//Отличие только в том, что данные в этом канале появятся не при явной записи в этот канал, а при вызове cancel() или
//функции метода ctx.Done() явно или, например, по истечению дедлайна.

func main() {
	var wg sync.WaitGroup
	wg.Add(6)

	//g0 will die after 3 sec, when read from done1
	done0 := make(chan struct{})
	go func() { time.Sleep(time.Second * 2); close(done0) }()

	//g1 will die after 3 sec, when read from done1
	done1 := make(chan struct{})
	go func() { time.Sleep(time.Second * 3); done1 <- struct{}{} }()

	//g2 will die after one of the signals, when read from done2
	done2 := make(chan os.Signal, 1)
	signal.Notify(done2, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	//g3 will die after 5 sec, when ctx3.Done() will be called and g3 will read from it
	ctx3, cancel3 := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel3()

	//g4 wil die after 7 sec, when ctx4.Done() will be called, and g4 will read from it
	ctx4, cancel4 := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel4()

	//g5 wil die after 5 sec, when cancel() will be called, and g5 will read from ctx.Done
	ctx5, cancel5 := context.WithTimeout(context.Background(), time.Second*120)
	go func() { time.Sleep(time.Second * 8); cancel5() }()

	go g0(done0, &wg)
	go g1(done1, &wg)
	go g2(done2, &wg)
	go g3(ctx3.Done(), &wg)
	go g4(ctx4, &wg)
	go g5(ctx5, &wg)

	wg.Wait()
	fmt.Println("All goroutines are dead")
}
func g0(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-done:
			fmt.Println("g0 is dead")
			return
		default:
			fmt.Println("I'm g0 and I'm alive!")
			time.Sleep(time.Millisecond * 1000)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func g1(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-done:
			fmt.Println("g1 is dead")
			return
		default:
			fmt.Println("I'm g1 and I'm alive!")
			time.Sleep(time.Millisecond * 1000)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func g2(done <-chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-done:
			fmt.Println("g2 is dead")
			return
		default:
			fmt.Println("I'm g2 and I'm alive!")
			time.Sleep(time.Millisecond * 1000)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func g3(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-done:
			fmt.Println("g3 is dead")
			return
		default:
			fmt.Println("I'm g3 and I'm alive!")
			time.Sleep(time.Millisecond * 1000)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func g4(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("g4 is dead")
			return
		default:
			fmt.Println("I'm g4 and I'm alive!")
			time.Sleep(time.Millisecond * 500)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

func g5(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("g5 is dead")
			return
		default:
			fmt.Println("I'm g5 and I'm alive!")
			time.Sleep(time.Millisecond * 1000)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}
