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
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//программа
	//ввод сообщения
	var source string
	fmt.Print("Введите слово: ")
	fmt.Scan(&source)

	//отправка сообщения серверу (всякий Read ожидает Write)
	if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
		fmt.Println(err)
		return
	}

	// получем ответ
	fmt.Print("клиент: сообщение отправлено")
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)

	if err != nil {
		return
	}

	//вывод
	fmt.Println()
	fmt.Print(string(buff[0:n]))
	fmt.Println()

	//всего 1 сообщение - требование 2го задания
	conn.Close()
}
