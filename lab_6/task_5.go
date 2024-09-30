package main

//5. Разработка многопоточного калькулятора:
//Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /).
//Используйте каналы для отправки запросов и возврата результатов.
//Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.

//что-то недосказано. я сделал по-факту горутины ввода и вывода, с какими-то дополнительными операциями
//в целом смогу доделать на паре, если скажут что конкретно нужно сделать

import (
	"fmt"
)

func task_5() {

	var un1 chan int = make(chan int)
	var un2 chan int = make(chan int)
	var us chan string = make(chan string)

	var fin chan bool = make(chan bool)

	go func() {
		fmt.Println("Введите число 1")
		var userNum1 int
		fmt.Scan(&userNum1)
		un1 <- userNum1

		fmt.Println("Введите число 2")
		var userNum2 int
		fmt.Scan(&userNum2)
		un2 <- userNum2

		for {
			fmt.Println("Введите операцию")
			var userStr string
			fmt.Scan(&userStr)

			if userStr == "+" || userStr == "-" || userStr == "*" || userStr == "/" {
				us <- userStr
			} else {
				fmt.Println("Необходим оператор, попытайтесь снова")
			}
		}
	}()

	//вывод
	go func() {
		n1 := <-un1
		n2 := <-un2

		sum := n1 + n2
		subtract := n1 - n2
		multiply := n1 * n2
		division := func() float64 {
			if n2 == 0 {
				return 0
			}
			return float64(n1) / float64(n2)
		}()

		switch <-us {
		case "+":
			fmt.Print("Результат сложения ")
			fmt.Println(sum)
		case "-":
			fmt.Print("Результат вычитания ")
			fmt.Println(subtract)
		case "*":
			fmt.Print("Результат умножения ")
			fmt.Println(multiply)
		case "/":
			fmt.Print("Результат деления ")
			fmt.Println(division)

		}
		fin <- true
	}()

	<-fin
	fmt.Println("Конец выполнения 5 задания")
}
