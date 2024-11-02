package main

// 5.	Тестирование API:
//   •	Реализуйте unit-тесты для каждого маршрута.
//   •	Проверьте корректность работы при различных вводных данных.

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func dbConnect() {
	connStr := "user=postgres password=123 dbname=go_lab8 sslmode=disable"
	db, _ = sql.Open("postgres", connStr)
}

func TestGet(t *testing.T) {
	//подключение к бд
	dbConnect()
	defer db.Close()

	//создание http-запроса
	r, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(usersHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Get Неверный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}
}

func TestGetOne(t *testing.T) {
	//подключение к бд
	dbConnect()
	defer db.Close()

	//создание http-запроса
	r, err := http.NewRequest("GET", "/users/80", nil)
	if err != nil {
		t.Fatal(err)
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		t.Errorf("GetOne Неверный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}
}

func TestPost(t *testing.T) {
	//подключение к бд
	dbConnect()
	defer db.Close()

	//создание данных
	userTest := User{Name: "Тест", Age: 50}
	data, _ := json.Marshal(userTest)

	//создание http-запроса
	r, err := http.NewRequest("POST", "/users", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(usersHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Post Неверный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}
}

func TestPut(t *testing.T) {
	//подключение к бд
	dbConnect()
	defer db.Close()

	//создание данных
	userTest := User{Name: "Тест2", Age: 25}
	data, _ := json.Marshal(userTest)

	//создание http-запроса
	r, err := http.NewRequest("PUT", "/users/78", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Put Неверный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}
}

func TestDelete(t *testing.T) {
	//подключение к бд
	dbConnect()
	defer db.Close()

	//создание http-запроса
	r, err := http.NewRequest("DELETE", "/users/79", nil)
	if err != nil {
		t.Fatal(err)
	}
	//тот кто возьмет ответ - httptest
	w := httptest.NewRecorder()

	//создание триггера и его ручной вызов
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(w, r)

	//проверка статуса запроса
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Delete Неверный код ответа: получили %v, ожидали %v", status, http.StatusOK)
	}
}
