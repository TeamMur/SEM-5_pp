package main

// 4.	Создание HTTP-сервера:
//Реализуйте базовый HTTP-сервер с обработкой простейших GET и POST запросов.
//Сервер должен поддерживать два пути:
//GET /hello — возвращает приветственное сообщение.
//POST /data — принимает данные в формате JSON и выводит их содержимое в консоль.

// 5.	Добавление маршрутизации и middleware:
//   •	Реализуйте обработку нескольких маршрутов и добавьте middleware для логирования входящих запросов.
//   •	Middleware должен логировать метод, URL, и время выполнения каждого запроса.

//NOTE: в коде нет обработки или вывода ошибок

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Name string `json:"name"` //WHY
	Age  int    `json:"age"`
}

var wg sync.WaitGroup

// Middleware обработчик-логгер НОВОЕ В 5 ЗАДАНИИ
func loggingMuddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
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
		json.NewDecoder(r.Body).Decode(&user)
		defer r.Body.Close() //WHY
		fmt.Println("Получено: ", user)
	}
}

func main() {
	//триггеры обернутые в логгер НОВОЕ В 5 ЗАДАНИИ
	http.Handle("/hello", loggingMuddleware(http.HandlerFunc(getHandler)))
	http.Handle("/data", loggingMuddleware(http.HandlerFunc(postHandler)))

	//запуск сервера !ATTENTION! блокирует поток до закрытия сервера поэтому отдельная горутина
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	//отправка запроса POST ! только после запуска сервера
	jsonStr := []byte(`{"name":"Тимур", "age":20}`)
	http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonStr)) //адрес, тип данных, данные в нужном формате

	//сервер будет работать бесконечно, пока мы вручную не закроем, т.к. нет ни одного wg.Done()
	wg.Add(1)
	wg.Wait()
}
