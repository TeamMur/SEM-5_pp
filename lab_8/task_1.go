package main

// 1.	Построение базового REST API:
//   •	Реализуйте сервер, поддерживающий маршруты:
//   •	GET /users — получение списка пользователей.
//   •	GET /users/{id} — получение информации о конкретном пользователе.
//   •	POST /users — добавление нового пользователя.
//   •	PUT /users/{id} — обновление информации о пользователе.
//   •	DELETE /users/{id} — удаление пользователя.

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{0, "Толик", 20},
	{1, "Миша", 19},
}

// users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	//GET /users — получение списка пользователей
	if r.Method == http.MethodGet {
		//отправка
		data := users
		json.NewEncoder(w).Encode(data)
	}
	//POST /users — добавление нового пользователя
	if r.Method == http.MethodPost {
		//получение данных (переменная user - то куда декодер поместит данные)
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.ID = len(users)

		//добавление элемента
		users = append(users, user)
	}
}

// конкретный user
func userHandler(w http.ResponseWriter, r *http.Request) {
	//GET /users/{id} — получение информации о конкретном пользователе
	if r.Method == http.MethodGet {
		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)
		//отправка
		data := users[id]
		json.NewEncoder(w).Encode(data)
	}
	//PUT /users/{id} — обновление информации о пользователе
	if r.Method == http.MethodPut {
		//получение данных (переменная user - то куда декодер поместит данные)
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)

		//добавление элемента
		user.ID = id
		users[id] = user
	}
	//DELETE /users/{id} — удаление пользователя
	if r.Method == http.MethodDelete {
		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)
		//удаление элемента
		// users = append(users[:id], users[id+1:]...)
		//логично именно аннулировать данные
		users[id] = User{id, "", 0}
	}
}

func main() {
	//триггеры
	http.Handle("/users", loggingMuddleware(http.HandlerFunc(usersHandler)))
	http.Handle("/users/{id}", loggingMuddleware(http.HandlerFunc(userHandler)))

	//запуск сервера
	fmt.Println("Сервер успешно запущен")
	http.ListenAndServe("localhost:8080", nil)

}

//мое удобство:

// логгер
func loggingMuddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

//тестирование через консоль и curl
//вводить в отдельной консольке, винда не принимает ' поэтому вместо них юзать варик с \"
// POST
//curl -X POST -H "Content-Type: application/json" -d '{"name": "Тимур", "age": 20}' http://localhost:8080/users
//curl -X POST -H "Content-Type: application/json" -d "{\"name\": \"Тимур\", \"age\": 20}" http://localhost:8080/users
// PUT
//curl -X PUT -H "Content-Type: application/json" -d '{"name": "Новый Типур", "age": 21}' http://localhost:8080/users/2
//curl -X PUT -H "Content-Type: application/json" -d "{\"name\": \"Новый Типур\", \"age\": 21}" http://localhost:8080/users/2
// DELETE
//curl -X DELETE http://localhost:8080/users/1
