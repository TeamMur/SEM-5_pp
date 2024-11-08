package main

// 4.	Реализация защищённого канала передачи данных (TLS):
//   •	Модифицируйте TCP-сервер и клиент из предыдущих лабораторных работ для работы через защищённый канал с использованием TLS.
//   •	Сервер должен поддерживать установку безопасного соединения, а клиент — проверять сертификат сервера перед отправкой данных.
//   •	Реализуйте взаимную аутентификацию на уровне сертификатов.

//сертификаты - публичные ключи, содержащии данные о владельце, о центре сертификации и т.п.

//ЗАПУСКАТЬ В ОТДЕЛЬНОЙ cmd перед клиентом

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	// загрузка сертификатов сервера сразу на tls
	serverCert, err := tls.LoadX509KeyPair("server_cert.pem", "server_key.pem")
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификатов сервера: %v", err)
	}

	// загрузка связанных сертификатов клиентов
	clientCertPool := x509.NewCertPool()
	clientCert, err := ioutil.ReadFile("client_cert.pem")
	if err != nil {
		log.Fatalf("Ошибка чтения сертификата клиента: %v", err)
	}
	clientCertPool.AppendCertsFromPEM(clientCert)

	// конфиг содержащий сертификат и "некоторые включенные галочки"
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    clientCertPool,        // связанные клиентские сертификаты
		ClientAuth:   tls.RequestClientCert, //проверка сертификата клиента
	}

	//создание сервера
	listener, err := tls.Listen("tcp", "localhost:8080", config)
	if err != nil {
		log.Fatalf("Ошибка создания сервера: %v", err)
	} else {
		fmt.Println("Сервер запущен")
	}
	defer listener.Close()

	//ожидание клиента
	fmt.Println("Ожидание подключения")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка при подключении:", err)
			continue
		} else {
			fmt.Println("Сертификат подтвержден. Успешное подключение")
		}
		//запуск подключения в отдельной горутине

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
