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
	var userInt, userInt2 int
	fmt.Println("4. Арифметические операции с двумя целыми числами:")
	fmt.Print("Введите число 1: ")
	fmt.Scan(&userInt)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&userInt2)

	sum := userInt + userInt2
	difference := userInt - userInt2
	multiply := userInt * userInt2
	var divide float64 = float64(userInt) / float64(userInt2)

	fmt.Println("Сумма ", sum)
	fmt.Println("Разность ", difference)
	fmt.Println("Произведение ", multiply)
	fmt.Print("Частное ")
	if userInt2 != 0 {
		fmt.Print(divide)
	} else {
		fmt.Print("ошибка - деление на ноль")
	}
	fmt.Println("\n")

	//Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой
	floatOperations()

	//Написать программу, которая вычисляет среднее значение трех чисел
	fmt.Println("6. Среднее значение 3х чисел:")
	var userFloat, userFloat2, userFloat3 float64
	fmt.Print("Введите число 1: ")
	fmt.Scan(&userFloat)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&userFloat2)
	fmt.Print("Введите число 3: ")
	fmt.Scan(&userFloat3)
	var resultSix = (userFloat + userFloat2 + userFloat3) / 3
	fmt.Printf("Среднее значение трёх чисел %0.2f\n", resultSix)
}

func floatOperations() {
	fmt.Println("5. Операция над числами с плавающей запятой:")
	var userA, userB float64
	fmt.Print("Введите число 1: ")
	fmt.Scan(&userA)
	fmt.Print("Введите число 2: \n")
	fmt.Scan(&userB)
	sum := userA + userB
	difference := userA - userB
	fmt.Printf("Сумма %0.2f\n", sum)
	fmt.Printf("Разность %0.2f\n\n", difference)
}
