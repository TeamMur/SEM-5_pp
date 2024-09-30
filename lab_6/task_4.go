package main

//4. Синхронизация с помощью мьютексов:
//Реализуйте программу, в которой несколько горутин увеличивают общую переменную-счётчик.
//Используйте мьютексы (sync.Mutex) для предотвращения гонки данных.
//Включите и выключите мьютексы, чтобы увидеть разницу в работе программы.

//важно что mutex.Lock() блокирует все горутины кроме первой. При этом конкретног порядка исполнения горутин нет, он случайный
//счетчик нарастает последовательно, потому что всякая случайная разблокированная горутина лишь прибавляет к счетчику 1

import (
	"fmt"
	"sync"
	"time"
)

func task_4() {

	var counter int = 0

	var ch chan bool = make(chan bool)

	var mutex sync.Mutex

	var iter int = 10
	fmt.Println("с синхронизацией:")
	for i := 0; i < iter; i++ {
		go func() {
			mutex.Lock()

			time.Sleep(time.Duration(i * i))
			counter++
			fmt.Println("с:", counter)
			mutex.Unlock()
			if counter == iter {
				ch <- true
			}
		}()
	}

	<-ch
	fmt.Println("без синхронизации:")
	counter = 0
	for i := 0; i < iter; i++ {
		go func() {
			time.Sleep(time.Duration(i * i))
			counter++
			fmt.Println("б:", counter)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Конец выполнения 4 задания")
}
