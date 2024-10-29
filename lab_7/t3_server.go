package main

// 3.	Асинхронная обработка клиентских соединений:
//Добавьте в сервер многопоточную обработку нескольких клиентских соединений.
//Используйте горутины для обработки каждого нового соединения.
//Реализуйте механизм graceful shutdown: сервер должен корректно завершать все активные соединения при остановке.

//каждый клиент способен отдавать лишь 1 сообщение (так требовало 2 задание)

//NOTE: в коде нет обработки или вывода ошибок

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//ввод данных
	var clientsMax int = 3
	var clientsTest int = 1
	var clientsCurrent int = 0

	fmt.Println("максимум ", clientsMax, " клиента")
	fmt.Println("Остановить сервер  на клиенте под номером:")
	fmt.Scan(&clientsTest)

	//создание сервера
	listener, _ := net.Listen("tcp", "localhost:8080")

	//ожидание клиента
	fmt.Println("Ожидание подключения")
	for {
		conn, _ := listener.Accept()

		wg.Add(1)
		go handleConnection(conn)

		clientsCurrent += 1
		fmt.Println("подключен клиент №", clientsCurrent)
		if clientsCurrent == clientsMax {
			break
		} else if clientsCurrent == clientsTest {
			//этот блок исключительно для теста остановки сервера
			//до того как вся его работа будет выполнена
			break

		}
	}

	wg.Wait() //ожидание завершения всех активных горутин
	listener.Close()
	fmt.Println("Все клиенты завершили работу\nСервер завершил работу")
}

// программа
func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer wg.Done()
	//прием данных
	input := make([]byte, (1024 * 4))
	n, err := conn.Read(input)
	if n == 0 {
		fmt.Println("подключение прервано")
	}
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
	}
	source := string(input[0:n])
	// вывод
	fmt.Println("сервер: получено - ", source)
	// отправка сообщения клиенту (всякий Read ожидает Write)
	conn.Write([]byte("сервер: новое сообщение"))
}
