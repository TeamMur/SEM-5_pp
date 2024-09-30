package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task_3() {
	var gch chan int = make(chan int)
	var pch chan string = make(chan string)

	go func() {
		fmt.Println("Генерирую")
		for i := 0; i < 10; i++ {
			gch <- rand.Intn(100)
		}
		close(gch)
	}()

	go func() {
		fmt.Println("Принтую")
		for n := range gch { //цикл работает до закрытия канала, итерация при новом n (т.к. <-ch прерывает работу горутины)
			if n%2 == 0 {
				pch <- "ч"
			} else {
				pch <- "н"
			}
		}
		close(pch)
	}()

	//for{} - while
	//switch блокирует горитину пока есть активные каналы, если получает данные - выполняет соотв. кейс
	//Если все каналы (от основного) закрыты, то select всё равно выберет один из кейсов
	//код в большинстве случаев сгенерирует чисел больше, чем определит их четность, т.к. генерация никогда не ставится на паузу в отличие от принта
	var endg bool = false
	var endp bool = false
Loop:
	for { //while
		select {
		case g, ok := <-gch:
			if !ok {
				endg = true
				break
			}
			fmt.Printf("число: %d\n ", g)
		case p, ok := <-pch:
			if !ok {
				endp = true
				break
			}
			fmt.Printf("четно: %s\n", p)
		}
		if endg && endp {
			break Loop
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Конец выполнения 3 задания")
}
