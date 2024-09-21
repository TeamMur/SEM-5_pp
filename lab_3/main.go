package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//1. Создать пакет mathutils с функцией для вычисления факториала числа.

	fmt.Println()
	//2. Использовать созданный пакет для вычисления факториала введенного пользователем числа.

	fmt.Println()
	//3. Создать пакет stringutils с функцией для переворота строки и использовать его в основной программе.

	fmt.Println()
	//4. Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.
	fmt.Println("4. Создание и заполнение массива:")
	var array [5]int

	for i := range &array {
		array[i] = rand.Intn(100)
	}
	fmt.Println(array)

	fmt.Println()
	//5. Создать срез из массива и выполнить операции добавления и удаления элементов.
	fmt.Println("5. Операции со срезом:")
	slice := array[:]
	fmt.Println(slice)

	new_num := 100
	slice = append(slice, new_num)
	id := 2
	slice = append(slice[:id], slice[id+1:]...)

	fmt.Printf("Добавлено число %d, удалено %d:", new_num, array[id])
	fmt.Println()
	fmt.Println(slice)

	fmt.Println()
	//6. Написать программу, которая создает срез из строк и находит самую длинную строку
	fmt.Println("6. Вывод самой длинной строки:")

	strings := []string{"Тимур", "Миша", "Толик"}
	fmt.Println("Срез строк:")
	fmt.Println(strings)

	var longest_string string
	for _, el := range strings {
		if len([]rune(longest_string)) < len([]rune(el)) {
			longest_string = el
		}
	}
	fmt.Printf("Длиннейшая строка: %s", longest_string)

	//
	fmt.Println()
	fmt.Println()
	fmt.Println("Конец выполнения программы")
}
