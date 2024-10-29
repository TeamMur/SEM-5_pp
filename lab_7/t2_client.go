package main

// 2.	Реализация TCP-клиента:
//Разработайте TCP-клиента, который подключается к вашему серверу.
//Клиент должен отправлять сообщение, введённое пользователем, и ожидать ответа.
//После получения ответа от сервера клиент завершает соединение.

//ЗАПУСКАТЬ В ОТДЕЛЬНОЙ cmd после сервера

import (
	"fmt"
	"net"
)

func main() {

	//подключение
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	//ввод сообщения
	var source string
	fmt.Print("Введите слово: ")
	fmt.Scan(&source)

	//отправка сообщения серверу (всякий Read ожидает Write)
	conn.Write([]byte(source))
	fmt.Println("клиент: сообщение отправлено")

	// получем ответ
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)

	//вывод
	fmt.Print(string(buff[0:n]))
}
