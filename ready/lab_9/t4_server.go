package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"

	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// User уже занято под таблицу
type Person struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var sessions = make(map[string]string)

// NOTE: массив вместо отдельной бд
var persons = make(map[string]string)

func generateToken() string {
	//генерация случайного числа числа
	b := make([]byte, 32)
	rand.Read(b)

	//преобразование в кодировку, в которой обычно представляется токен
	//проще говоря это "нужный тип данных" для токена - ключа-доступа к чему-либо
	return base64.URLEncoding.EncodeToString(b)
}

func authorizeUser(w http.ResponseWriter, r *http.Request) {
	//декодирование пользователя из post-запроса
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	//проверка логина
	if _, exists := persons[person.Login]; !exists {
		http.Error(w, "Пользователя не существует", http.StatusUnauthorized)
		return
	}
	//проверка пароля
	if person.Password != persons[person.Login] {
		http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		return
	}

	//генерация токена
	token := generateToken()

	//установка токена
	sessions[token] = person.Login
	w.Header().Set("Authorization", token)

	fmt.Println("Новое подключение: \n", w.Header().Get("Authorization"))

	//установка статуса (успешно)
	w.WriteHeader(http.StatusOK)
}

// оболочка функций для проверки авторизованности при каждом запроса
func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//получаем токен с реквеста
		token := r.Header.Get("Authorization")
		//проверяем на наличие в массиве
		if _, exists := sessions[token]; !exists {
			http.Error(w, "Неавторизован", http.StatusUnauthorized)
			return
		}
		//выполняем запрос
		next.ServeHTTP(w, r)
	})
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres password=123 dbname=go_lab8 sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	defer db.Close()
	//GET /users — получение списка пользователей
	if r.Method == http.MethodGet {
		//получение переменных из запроса
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		nameF := r.URL.Query().Get("name")
		ageF, _ := strconv.Atoi(r.URL.Query().Get("age"))

		//формирование запроса (с валидацией)
		query := "SELECT * from users"
		if nameF != "" || ageF > 0 {
			query += " WHERE "
			if nameF != "" {
				query += "name = '" + nameF + "'"
			}
			if nameF != "" && ageF > 0 {
				query += " AND "
			}
			if ageF > 0 {
				query += "age = " + strconv.Itoa(ageF)
			}
		}
		if limit > 0 {
			query += " LIMIT " + strconv.Itoa(limit)
		}

		if offset > 0 {
			query += " OFFSET " + strconv.Itoa(offset)
		}
		fmt.Println("Выполняется: ", query)
		//получение всех строк
		rows, err := db.Query(query)
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
	connStr := "user=postgres password=123 dbname=go_lab8 sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	defer db.Close()
	//GET /users/{id} — получение информации о конкретном пользователе
	if r.Method == http.MethodGet {
		//получение id и конвертация в int
		base := path.Base(r.URL.Path)
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
		base := path.Base(r.URL.Path)
		id, _ := strconv.Atoi(base)

		//обновление
		user.ID = id
		db.Exec("UPDATE users SET name = $2, age = $3 WHERE id = $1", user.ID, user.Name, user.Age)
	}
	//DELETE /users/{id} — удаление пользователя
	if r.Method == http.MethodDelete {
		//получение id и конвертация в int
		base := path.Base(r.URL.Path)
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

// мое удобство:
// логгер
func loggingMuddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
