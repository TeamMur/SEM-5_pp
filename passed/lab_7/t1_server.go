package main

// 1.	Создание TCP-сервера:
//Реализуйте простой TCP-сервер, который слушает указанный порт и принимает входящие соединения.
//Сервер должен считывать сообщения от клиента и выводить их на экран.
//По завершении работы клиенту отправляется ответ с подтверждением получения сообщения.

//ЗАПУСКАТЬ В ОТДЕЛЬНОЙ cmd перед клиентом

//NOTE: в коде нет обработки или вывода ошибок

import (
	"fmt"
	"net"
)

func main() {
	//создание сервера
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()

	//ожидание клиента
	fmt.Println("Ожидание подключения")
	for {
		conn, _ := listener.Accept()
		//запуск подключения в отдельной горутине
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
	//вывод
	source := string(input[0:n])
	fmt.Println("сервер: слово - ", source)
	//отправка сообщения клиенту (всякий Read ожидает Write)
	conn.Write([]byte("сервер: новое сообщение"))
}
