package main

import (
	"database/sql"
	"fmt"
	"strings"

	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

var roleToken string
var roleUser string

var jwtKey = []byte("my_secret_key")

// генерация токена на основе роли пользователя
func generateJWT(role string) (string, error) {
	claims := &jwt.StandardClaims{Subject: role}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// проверка токена перед выполнением запроса
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//проверка наличия
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			fmt.Println("токен отсутствует")
			return
		}

		// Проверяем, что токен начинается с "Bearer "
		fields := strings.Fields(tokenString)
		if len(fields) != 2 || fields[0] != "Bearer" {
			fmt.Println("Неверный формат токена")
			return
		}

		tokenStr := fields[1]
		if tokenStr != roleToken {
			fmt.Println("Неверный токен")
			return
		}

		// Парсим и проверяем токен
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Неверный или просроченный токен")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Метод не поддерживается")
		return
	}

	// то куда поместиться инфа из запроса
	var loginReq struct {
		Role string `json:"role"`
	}

	// получение инфы из запроса
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil || loginReq.Role == "" {
		fmt.Println("Неверный формат запроса")
		return
	}

	// генерация токена на основе роли
	token, err := generateJWT(loginReq.Role)
	if err != nil {
		fmt.Println("Ошибка генерации токена")
		return
	}

	roleUser = loginReq.Role
	roleToken = token
	fmt.Println("Ваша роль изменена на ", roleUser)
	fmt.Println("Ваш токен: ", roleToken)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Выполнение запроса...")
	//GET /users — получение списка пользователей
	if r.Method == http.MethodGet {
		//получение переменных из запроса
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		nameF := r.URL.Query().Get("name")
		ageF, err := strconv.Atoi(r.URL.Query().Get("age"))
		//обработка ошибок
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
		return
	}

	if roleUser != "admin" {
		fmt.Println("Недостаточно прав")
		return
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
	fmt.Println("Выполнение запроса...")
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
		return
	}
	if roleUser != "admin" {
		fmt.Println("Недостаточно прав")
		return
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

var db *sql.DB

func main() {
	//подключение к бд
	connStr := "user=postgres password=123 dbname=go_lab8 sslmode=disable"
	db, _ = sql.Open("postgres", connStr)

	//триггеры
	http.Handle("/users", loggingMuddleware(authMiddleware(http.HandlerFunc(usersHandler))))
	http.Handle("/users/{id}", loggingMuddleware(authMiddleware(http.HandlerFunc(userHandler))))
	http.HandleFunc("/login", loginHandler)

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
		log.Printf("Request: %s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

//получить токен для переданной роли
//curl -X POST -H "Authorization: Bearer <ваш_токен>" -H "Content-Type: application/json" -d "{\"role\": \"admin\"}" http://localhost:8080/login

//тестирование через консоль и curl
// POST
//curl -X POST -H "Authorization: Bearer <ваш_токен>" -H "Content-Type: application/json" -d "{\"name\": \"Тимур\", \"age\": 20}" http://localhost:8080/users
// PUT
//curl -X PUT -H "Authorization: Bearer <ваш_токен>" -H "Content-Type: application/json" -d "{\"name\": \"Бахром\", \"age\": 20}" http://localhost:8080/users/2
// DELETE
//curl -X DELETE -H "Authorization: Bearer <ваш_токен>" http://localhost:8080/users/1
