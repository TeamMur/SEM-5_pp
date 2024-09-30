package main

//1. Создание и запуск горутин
//Напишите программу, которая параллельно выполняет три функции
//Каждая функция должна выполняться в своей горутине.
//Добавьте использование time.Sleep() для имитации задержек и продемонстрируйте параллельное выполнение.

import (
	"fmt"
	"math/rand"
	"time"
)

func task_1() {
	for i := range 10 {
		go factorial(i)
		go random(i)
		go justn(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Конец выполнения 1 задания")
}

// расчёт факториала
func factorial(n int) {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("fact:", n, "-", result)
}

// генерация случайных чисел
func random(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("rand:", n, "-", rand.Intn(100))
}

// 3 функция
func justn(n int) {
	var result int = 0
	for i := range n {
		for j := range n {
			result += i + j
		}
	}
	time.Sleep(150 * time.Millisecond)
	fmt.Println("just:", n, "-", result)
}
