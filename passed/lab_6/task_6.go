package main

//6. Создание пула воркеров:
//Реализуйте пул воркеров, обрабатывающих задачи (например, чтение строк из файла и их реверсирование).
//Количество воркеров задаётся пользователем.
//Распределение задач и сбор результатов осуществляется через каналы.
//Выведите результаты работы воркеров в итоговый файл или в консоль.

import (
	"fmt"
)

func worker(tasks <-chan int, results chan<- data, chars []rune) {
	for t := range tasks { //итерация только при получении значения, при close выход из цикла
		ri := len(chars) - t - 1
		results <- data{ri, chars[ri]}

	}
}

func task_6() {

	var start_string string = "Здаров"
	var start_chars []rune = []rune(start_string)
	var final_chars []rune = []rune(start_string)
	var size int = len(start_chars)

	fmt.Println(size)

	//6.1 Ввод количества воркеров
	fmt.Println("6. Введите количество воркеров")
	var userNum1 int
	fmt.Scan(&userNum1)

	//Каналы
	//tasks для передачи индексов строкового массива
	//results для передачи data: id - индекс символа, char - символ
	tasks := make(chan int, size)
	results := make(chan data, size)

	//Запуск воркеров: они будут ожидать ввода любых данных
	for w := 0; w < userNum1; w++ {
		go worker(tasks, results, start_chars)
	}

	//Передача данных воркерам
	for t := 0; t < size; t++ {
		tasks <- t
	}

	//Закрытие канала tasks, чтобы каждый воркер не продолжал цикл
	close(tasks)

	//Получение результатов, чтобы каждый воркер смог прекратить итерацию
	for a := 0; a < size; a++ {
		d := <-results
		i := size - d.id - 1
		final_chars[i] = d.char
	}

	//6.4. Вывод результатов
	var final_string string = string(final_chars)
	fmt.Println(final_string)
}

type data struct {
	id   int
	char rune
}
