package main

// 6.	Веб-сокеты:
//   •	Реализуйте сервер на основе веб-сокетов для чата.
//   •	Клиенты должны подключаться к серверу, отправлять и получать сообщения.
//   •	Сервер должен поддерживать несколько клиентов и рассылать им сообщения, отправленные любым подключённым клиентом.

//веб-сокет это протокол, т.е. особый формат передачи данных
//веб-сокет предполагает двустороннюю коммуникацию, где сообщения могут исходить как от клиента, так и от сервера

//сервер переводиться в вебсокет режим. значит и подключение должно быть в вебсокет режиме, можно использовать сайт
//https://piehost.com/websocket-tester    где нужно вставить ссылку     ws://localhost:8080/ws

//NOTE: в коде нет обработки или вывода ошибок

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// карта (массив) подключений
var clients = make(map[*websocket.Conn]bool)

// канал передачи текста (сообщений)
var mChan = make(chan string)

// преобразователь http-соединения в websocket-соединение
var upgrader = websocket.Upgrader{
	//??
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Обработчик веб-сокет соединений
func handleConnections(w http.ResponseWriter, r *http.Request) {
	//преобразуем соединение
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	//добавляем клиента в карту (массив с ключами)
	clients[ws] = true

	for {
		//чтение сообщения (type, string, err: type и err опущены)
		_, msg, _ := ws.ReadMessage()
		//размещение сообщения в канал
		mChan <- string(msg)
	}
}

// рассылка сообщений
func handleMessages() {
	//бесконечный цикл с итерацией лишь при каждом получении строки в канал
	for {
		//получение сообщения
		msg := <-mChan
		fmt.Println("Запрошена отправка: ", msg)

		//отправка сообщения всем клиентам
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
}

func main() {
	//запуск обработчика сообщений
	go handleMessages()

	//триггер на ../ws странице: переданный метод работает для каждого подключившегося
	http.HandleFunc("/ws", handleConnections)

	//запуск сервера
	fmt.Println("Сервер успешно запущен")
	http.ListenAndServe("localhost:8080", nil)
}
