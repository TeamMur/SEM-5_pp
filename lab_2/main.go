package main

import (
	"fmt"
)

func main() {
	//1. Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
	fmt.Println("Определение четности числа:")
	fmt.Print("Введите число: ")
	var userInt int
	fmt.Scan(&userInt)
	if userInt%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

	fmt.Println()
	//2. Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".
	fmt.Println("Определение сигнатуры числа:")
	fmt.Print("Введите число: ")

	var userFloat float64
	fmt.Scan(&userFloat)
	if userFloat > 0 {
		fmt.Println("Positive")
	} else if userFloat < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	fmt.Println()
	//3. Написать программу, которая выводит все числа от 1 до 10 с помощью цикла for.
	fmt.Println("Числа от 1 до 10:")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("\n")
	//4. Написать функцию, которая принимает строку и возвращает ее длину.
	fmt.Println("Определение длины строки")
	fmt.Print("Введите строку: ")

	var userStr string
	fmt.Scan(&userStr)
	fmt.Println("Длина строки '", userStr, "':")
	fmt.Println(get_string_length(userStr))

	fmt.Println()
	//5. Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
	fmt.Println("Расчет площади прямоугольника:")

	var userRectA, userRectB float64
	fmt.Print("Введите высоту: ")
	fmt.Scan(&userRectA)
	fmt.Print("Введите ширину: ")
	fmt.Scan(&userRectB)

	rect := Rectangle{userRectA, userRectB}
	fmt.Println(get_rectangle_area(rect))

	fmt.Println()
	//6. Написать функцию, которая принимает два целых числа и возвращает их среднее значение.
	fmt.Println("Среднее двух целых чисел:")

	var userInt2 int
	fmt.Print("Число 1: ")
	fmt.Scan(&userInt)
	fmt.Print("Число 2: ")
	fmt.Scan(&userInt2)

	fmt.Println(two_nums_average(userInt, userInt2))
}

// 4.
func get_string_length(str string) int {
	return len([]rune(str))
}

// 5.
type Rectangle struct {
	a, b float64
}

func get_rectangle_area(rect Rectangle) float64 {
	return rect.a * rect.b
}

// 6.
func two_nums_average(a int, b int) float64 {
	return float64(a+b) / 2
}
