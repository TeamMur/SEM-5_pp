package main

// 2.	Подключение базы данных:
//   •	Добавьте базу данных (например, PostgreSQL или MongoDB) для хранения информации о пользователях.
//   •	Модифицируйте сервер для взаимодействия с базой данных.

import (
	"database/sql"
	"fmt"

	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	//GET /users — получение списка пользователей
	if r.Method == http.MethodGet {
		//получение всех строк
		rows, err := db.Query("SELECT * from users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		//извлечение данных
		result := []User{}
		for rows.Next() {
			user := User{}
			rows.Scan(&user.ID, &user.Name, &user.Age)
			result = append(result, user)
		}

		//отправка
		data := result
		json.NewEncoder(w).Encode(data)
	}
	//POST /users — добавление нового пользователя
	if r.Method == http.MethodPost {
		//получение данных (переменная user - то куда декодер поместит данные)
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		//валидация и обработка ошибок
		if err != nil || user.Name == "" || user.Age <= 0 {
			http.Error(w, "Неверный формат или невозможные значения данных", http.StatusBadRequest)
			return
		}

		//добавление элемента
		db.Exec("INSERT INTO users (name, age) values ($1, $2)", user.Name, user.Age)
	}
}

// конкретный user
func userHandler(w http.ResponseWriter, r *http.Request) {
	//GET /users/{id} — получение информации о конкретном пользователе
	if r.Method == http.MethodGet {
		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)
		//Выборка
		var user User
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//отправка
		data := user
		json.NewEncoder(w).Encode(data)
	}
	//PUT /users/{id} — обновление информации о пользователе
	if r.Method == http.MethodPut {
		//получение данных (переменная user - то куда декодер поместит данные)
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		//валидация и обработка ошибок
		if err != nil || user.Name == "" || user.Age <= 0 {
			http.Error(w, "Неверный формат или невозможные значения данных", http.StatusBadRequest)
			return
		}

		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)

		//обновление
		user.ID = id
		db.Exec("UPDATE users SET name = $2, age = $3 WHERE id = $1", user.ID, user.Name, user.Age)
	}
	//DELETE /users/{id} — удаление пользователя
	if r.Method == http.MethodDelete {
		//получение id и конвертация в int
		base := path.Base(r.RequestURI)
		id, _ := strconv.Atoi(base)

		//Удаление
		_, err := db.Exec("DELETE FROM users WHERE id = $1", id)

		//обработка ошибок
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

var db *sql.DB

func main() {
	//подключение к бд
	connStr := "user=postgres password=123 dbname=go_lab8 sslmode=disable"
	db, _ = sql.Open("postgres", connStr)

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
//curl -X PUT -H "Content-Type: application/json" -d '{"name": "Бахром", "age": 20}' http://localhost:8080/users/2
//curl -X PUT -H "Content-Type: application/json" -d "{\"name\": \"Бахром\", \"age\": 20}" http://localhost:8080/users/2
// DELETE
//curl -X DELETE http://localhost:8080/users/1
