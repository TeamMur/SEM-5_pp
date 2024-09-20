package main

import (
	"fmt"
	"strings"
)

func main() {
	//1. Написать программу, которая создает карту с именами людей и их возрастами. Добавить нового человека и вывести все записи на экран.
	fmt.Println("1. Карта с именами и возрастами:")

	fmt.Println("Карта до изменений:")
	persons := map[string]int{"Тимур": 20, "Миша": 19}
	fmt.Println(persons)

	fmt.Println("Карта после изменений:")
	persons["Толик"] = 19
	fmt.Println(persons)

	fmt.Println()
	//2. Реализовать функцию, которая принимает карту и возвращает средний возраст всех людей в карте
	fmt.Println("2. Средний возраст всех людей в карте:")
	result_second := get_aver_age(persons)
	fmt.Println(result_second)

	fmt.Println()
	//3. Написать программу, которая удаляет запись из карты по заданному имени.
	fmt.Println("3. Удаление записи из карты:")

	fmt.Println("Карта до изменений:")
	fmt.Println(persons)

	fmt.Println("Карта после изменений:")
	delete(persons, "Миша")
	fmt.Println(persons)

	fmt.Println()
	//4. Написать программу, которая считывает строку с ввода и выводит её в верхнем регистре.
	fmt.Println("4. Вывод строки в верхнем регистре:")

	fmt.Println("Введите строку:")
	var userStr string
	fmt.Scan(&userStr)

	userStr = strings.ToUpper(userStr)
	fmt.Println(userStr)

	fmt.Println()
	//5. Написать программу, которая считывает несколько чисел, введенных пользователем, и выводит их сумму.
	fmt.Println("5. Сумма нескольких чисел:")
	fmt.Println("Введите число чисел:")
	var user_num_count int
	fmt.Scan(&user_num_count)

	var sum float64
	for i := 0; i < user_num_count; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		var new_num float64
		fmt.Scan(&new_num)
		sum += new_num
	}
	fmt.Printf("Сумма: %0.2f", sum)

	fmt.Println()
	//6. Написать программу, которая считывает массив целых чисел и выводит их в обратном порядке.
	fmt.Println("6. Массив в обратном порядке:")
	fmt.Println("Введите размер массива:")
	fmt.Scan(&user_num_count)
	var nums []float64
	for i := 0; i < user_num_count; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		var new_num float64
		fmt.Scan(&new_num)
		nums = append(nums, new_num)
	}
	fmt.Println("Ваш массив:")
	fmt.Println(nums)
	fmt.Println("Вывод массива в обратном порядке:")
	for i := user_num_count - 1; i >= 0; i-- {
		fmt.Print(nums[i], " ")
	}

	fmt.Println("\nКонец выполнения программы")
}

// 2.
func get_aver_age(persons map[string]int) float64 {
	var age_sum int
	for key := range persons {
		age_sum += persons[key]
	}
	return float64(age_sum) / float64(len(persons))
}
