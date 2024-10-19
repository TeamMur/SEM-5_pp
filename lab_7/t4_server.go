package main

// 4.	Создание HTTP-сервера:
//Реализуйте базовый HTTP-сервер с обработкой простейших GET и POST запросов.
//Сервер должен поддерживать два пути:
//GET /hello — возвращает приветственное сообщение.
//POST /data — принимает данные в формате JSON и выводит их содержимое в консоль.

//ЗАПУСКАЙ В БРАУЗЕРЕ АДРЕС "localhost:8080/папка"
//параметры в запросе:
//url&name=value&name2=value2
//example.com?name=value

//ЗАПУСКАЙ ИЗ КОНСОЛИ (ДЛЯ КОНКРЕТНЫХ ЗАПРОСОВ)
//curl -X POST -d "name=value" url

//request - запрос
//responce - ответ

//HandleFunc
//первый параметр - папка в которой будет обработка
//второй - функция - ответ на любой запрос в указанной папке

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"` //WHY
	Age  int    `json:"age"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	//w - интерфейс для отправки ответа клиенту
	//r - запрос
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name") //получить значение по ключу в url, его может и не быть, тогда значение = "" (string)
		if name != "" {
			fmt.Fprintf(w, "Привет, %s!", name)
		} else {
			fmt.Fprintln(w, "Привет, мир!")
		}

	}
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Ошибка при декодировании")
			return
		}
		defer r.Body.Close() //WHY
		fmt.Println("Получено: ", user)
	}
}

func main() {
	//триггеры
	http.HandleFunc("/hello", getHandler)
	http.HandleFunc("/data", postHandler)

	//запуск сервера !ATTENTION! блокирует поток до закрытия сервера поэтому отдельная горутина
	go func() {
		err2 := http.ListenAndServe("localhost:8080", nil)
		if err2 != nil {
			fmt.Println("Ошибка при запуске сервера:", err2)
		}
	}()

	//отправка запроса POST ! только после запуска сервера
	jsonStr := []byte(`{"name":"Тимур", "age":20}`)
	resp, err := http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonStr)) //адрес, тип данных, данные в нужном формате
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //WHY предполагаю прекратить получение данных из запроса, чтобы не занимало оперативку, т.к. возможно это штука не самоочищается т.к. вообще другой яп

}
