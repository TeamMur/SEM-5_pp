package main

import (
	"fmt"
	"time"
)

func main() {
	//Вывод текущей даты и времени
	var timeNow = time.Now()
	fmt.Printf("\n1. Текущая дата и время:\n%v\n", timeNow.Format(time.DateTime))

	/*Создать переменные различных типов (int, float64, string, bool) и вывести их на экран
	Использовать краткую форму объявления переменных для создания и вывода переменных*/
	varInt := 5
	varFloat64 := 3.14
	varString := "Тимур"
	varBool := true
	fmt.Printf("\n2-3. Переменные различных типов:\nint %v\nfloat64 %v\nstring %v\nbool %v\n\n", varInt, varFloat64, varString, varBool)

	//Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов
	num1 := 2
	num2 := 3
	fmt.Println("4. Арифметические операции с двумя целыми числами:")
	fmt.Println("Сумма ", num1+num2)
	fmt.Println("Разность ", num1-num2)
	fmt.Println("Произведение ", num1*num2)
	fmt.Print("Частное ")
	if num2 != 0 {
		fmt.Print(num1 / num2)
	} else {
		fmt.Print("ошибка")
	}
	fmt.Println("\n")

	//Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой
	floatOperations(1.2, 4.3)

	//Написать программу, которая вычисляет среднее значение трех чисел
	avNum1 := 2
	avNum2 := 4.23
	avNum3 := 5.3
	fmt.Printf("Среднее значение трёх чисел %0.2f\n", (float64(avNum1)+avNum2+avNum3)/3)
}

func floatOperations(a float32, b float32) {
	fmt.Println("Операция над числами с плавающей запятой:")
	fmt.Printf("Сумма %0.2f\n", a+b)
	fmt.Printf("Разность %0.2f\n\n", a-b)
}
