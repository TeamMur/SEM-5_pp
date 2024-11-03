package main

// 1.	Создание клиентской программы:
//   •	Напишите клиентское приложение, которое отправляет запросы на ваш сервер.
//   •	Реализуйте простое меню с возможностью выполнения CRUD операций.

// 2.	Обработка ответов:
//   •	Реализуйте механизм обработки ответов и ошибок.
//   •	Добавьте функции для форматирования и вывода полученных данных на экран.

// 3.	Интерфейс пользователя:
//   •	Разработайте консольный интерфейс с меню для пользователя.
//   •	Включите возможность добавления, удаления и обновления информации о пользователях.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

func main() {
	//Консольное приложение
	//ВЫБОР ОПЕРАЦИИ
	fmt.Println("Приветствую вас! Эта программа предназначена для работы с базой данных")

	const (
		sPost   string = "1. POST - создать данные"
		sGet    string = "2. GET - вывести существующие данные"
		sPut    string = "3. PUT - обновить существующие данные"
		sDelete string = "4. DELETE - удалить существующие данные"
		sExit   string = "5. Выход"
	)
	sOperations := [...]string{sPost, sGet, sPut, sDelete, sExit}
	var operationNum int

	for {
		fmt.Printf("Выберите операцию: \n%s\n%s\n%s\n%s\n%s\n", sPost, sGet, sPut, sDelete, sExit)
		fmt.Println("Ожидание ввода данных (требуется номер операции):")
		for {
			fmt.Scan(&operationNum)
			if operationNum < 1 || operationNum > 5 {
				fmt.Println("Попытайтесь снова. Число должно быть в пределах от 1 до 5")
			} else {
				break
			}
		}

		var operationString string = sOperations[operationNum-1]
		fmt.Println(operationString)
		switch operationNum {
		case 1: //POST
			var nUser User //new user
			var nName string
			var nAge int
			fmt.Println("Введите имя нового пользователя:")
			fmt.Scan(&nName)
			fmt.Println("Введите возраст нового пользователя:")
			for {
				fmt.Scan(&nAge)
				if nAge < 1 || nAge > 100 {
					fmt.Println("Попытайтесь снова. Число должно быть в пределах от 1 до 100")
				} else {
					break
				}
			}
			nUser.Name = nName
			nUser.Age = nAge
			Post(nName, nAge)
		case 2: //GET
			fmt.Println("Выборка данных. Возможные операции:\n1. Вывести все данные\n2. Вывести данные по id")
			var operationNum int
			for {
				fmt.Scan(&operationNum)
				if operationNum < 1 || operationNum > 2 {
					fmt.Println("Попытайтесь снова. Число должно быть в пределах от 1 до 2")
				} else {
					break
				}
			}
			if operationNum == 1 {
				Get()
			} else {
				//вывод конкретного
				fmt.Println("Введите существующий id пользователя")
				var inID int //input id
				fmt.Scan(&inID)
				GetOne(inID)
			}
		case 3: //PUT
			var inID int //input id
			var nName string
			var nAge int
			fmt.Println("Введите существующий id пользователя")
			fmt.Scan(&inID)

			fmt.Println("Введите имя нового пользователя:")
			fmt.Scan(&nName)
			fmt.Println("Введите возраст нового пользователя:")
			for {
				fmt.Scan(&nAge)
				if nAge < 1 || nAge > 100 {
					fmt.Println("Попытайтесь снова. Число должно быть в пределах от 1 до 100")
				} else {
					break
				}
			}
			Put(inID, nName, nAge)
		case 4: //DELETE
			fmt.Println("Введите существующий id пользователя")
			var inID int //input id
			fmt.Scan(&inID)
			Delete(inID)
		case 5: //Exit
			fmt.Println("Конец выполнения программы")
			return
		}
	}
}

// ОПЕРАЦИИ
func Get() {
	//создание http-запроса
	r, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(usersHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		fmt.Printf("Get Неверный код ответа: получили %v, ожидали %v\n", status, http.StatusOK)
		return
	}

	fmt.Println("Ответ:")
	printWBody(w.Body)
}

func GetOne(id int) {
	//создание http-запроса
	nPath := "/users/" + strconv.Itoa(id)
	r, err := http.NewRequest("GET", nPath, nil)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		fmt.Printf("GetOne Неверный код ответа: получили %v, ожидали %v\n", status, http.StatusOK)
		fmt.Println("Вероятно, вы указали несуществующий id")
		return
	}

	fmt.Println("Ответ:")
	printWBody(w.Body)
}

func Post(name string, age int) {
	//создание данных
	nUser := User{Name: name, Age: age}
	data, _ := json.Marshal(nUser)

	//создание http-запроса
	r, err := http.NewRequest("POST", "/users", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(usersHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		fmt.Printf("Post Неверный код ответа: получили %v, ожидали %v\n", status, http.StatusOK)
		return
	}

	fmt.Println("Ответ: \nПользователь успешно создан!")
}

func Put(id int, name string, age int) {
	//создание данных
	nUser := User{Name: name, Age: age}
	data, _ := json.Marshal(nUser)

	//создание http-запроса
	nPath := "/users/" + strconv.Itoa(id)
	r, err := http.NewRequest("PUT", nPath, bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		fmt.Printf("Put Неверный код ответа: получили %v, ожидали %v\n", status, http.StatusOK)
		fmt.Println("Вероятно, вы указали несуществующий id")
		return
	}

	fmt.Printf("Ответ: \nПользователь с id = %d успешно обновлен!\n", id)
}

func Delete(id int) {
	//создание http-запроса
	nPath := "/users/" + strconv.Itoa(id)
	r, err := http.NewRequest("DELETE", nPath, nil)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		fmt.Printf("Delete Неверный код ответа: получили %v, ожидали %v\n", status, http.StatusOK)
		fmt.Println("Вероятно, вы указали несуществующий id")
		return
	}

	fmt.Printf("Ответ: \nПользователь с id = %d удален\n", id)
}

// функция форматирования вывода w.Body
func printWBody(w *bytes.Buffer) {
	nStr := w.String()
	nStr = nStr[2:]
	nStr = strings.ReplaceAll(nStr, "[", "")
	nStr = strings.ReplaceAll(nStr, "]", "")
	nStr = strings.ReplaceAll(nStr, "{", "\n")
	nStr = strings.ReplaceAll(nStr, "}", "")
	nStr = strings.ReplaceAll(nStr, ":", "")
	nStr = strings.ReplaceAll(nStr, ",", " ")
	nStr = strings.ReplaceAll(nStr, "\"", "")
	nStr = strings.ReplaceAll(nStr, "id", "")
	nStr = strings.ReplaceAll(nStr, "name", "")
	nStr = strings.ReplaceAll(nStr, "age", "\t")

	split := strings.Split(nStr, "\n")
	fmt.Println("id name \tage")
	for _, s := range split {
		fmt.Println(s)
	}
}
