package main

// 1.	Создание TCP-сервера:
//Реализуйте простой TCP-сервер, который слушает указанный порт и принимает входящие соединения.
//Сервер должен считывать сообщения от клиента и выводить их на экран.
//По завершении работы клиенту отправляется ответ с подтверждением получения сообщения.

//ЗАПУСКАТЬ В ОТДЕЛЬНОЙ cmd перед клиентом

import (
	"fmt"
	"net"
)

func main() {
	//создание сервера
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//ожидание клиента
	fmt.Println("Ожидание подключения")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		fmt.Println("подключено")
		go handleConnection(conn)
	}
}

// программа
func handleConnection(conn net.Conn) {
	defer conn.Close()
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
	fmt.Println("сервер: слово - ", source)
	// отправка сообщения клиенту (всякий Read ожидает Write)
	conn.Write([]byte("сервер: новое сообщение"))
}
