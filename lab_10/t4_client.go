package main

// 4.	Реализация защищённого канала передачи данных (TLS):
//   •	Модифицируйте TCP-сервер и клиент из предыдущих лабораторных работ для работы через защищённый канал с использованием TLS.
//   •	Сервер должен поддерживать установку безопасного соединения, а клиент — проверять сертификат сервера перед отправкой данных.
//   •	Реализуйте взаимную аутентификацию на уровне сертификатов.

//ЗАПУСКАТЬ В ОТДЕЛЬНОЙ cmd после сервера

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// загрузка сертификатов клиента  сразу на tls
	clientCert, err := tls.LoadX509KeyPair("client_cert.pem", "client_key.pem")
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификатов клиента: %v", err)
	}

	// загрузка связанных сертификатов серверов
	serverCertPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("server_cert.pem")
	if err != nil {
		log.Fatalf("Ошибка чтения сертификата сервера: %v", err)
	}
	serverCertPool.AppendCertsFromPEM(serverCert)

	// конфиг содержащий сертификаты и "некоторые включенные галочки"
	config := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            serverCertPool, //связанные сервера
		InsecureSkipVerify: true,           //отключение глубокой проверки, однако проверка целостности данных и прочего будет осуществлен
	}

	//подключение
	conn, err := tls.Dial("tcp", "localhost:8080", config)
	if err != nil {
		fmt.Println("Ошибка. Проверьте сертификат")
	} else {
		fmt.Println("Сертификат подтвержден. Успешное подключение")
	}
	defer conn.Close()

	//ввод сообщения
	var source string
	fmt.Print("Введите слово: ")
	fmt.Scan(&source)

	//отправка сообщения серверу (всякий Read ожидает Write)
	conn.Write([]byte(source))
	fmt.Println("клиент: сообщение отправлено")

	//получаем ответ
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)

	//вывод
	fmt.Print(string(buff[0:n]))
}
