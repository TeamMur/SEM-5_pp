package main

// 3.	Асимметричное шифрование и цифровая подпись:
//   •	Создайте пару ключей (открытый и закрытый) и сохраните их в файл.
//   •	Реализуйте программу, которая подписывает сообщение с помощью закрытого ключа и проверяет подпись с использованием открытого ключа.
//   •	Продемонстрируйте пример передачи подписанных сообщений между двумя сторонами.

//Асимметричное шифрование - это шифрование с двумя ключами. Первый известен всем, а второй приватный
//ключи математически связаны, публичный ключ исходит от приватного. публичный шифрует, приватный дешифрует
//pem - методы и формат передачи таких ключей

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateKeys() {
	//ПРИВАТНЫЙ КЛЮЧ
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	//файл с приватным ключом
	privFile, _ := os.Create("private_key.pem")
	defer privFile.Close()

	// преобразование ключа в стандартизированный вид передачи
	privBytes := x509.MarshalPKCS1PrivateKey(privKey)

	// маркировка (-BEGIN- -END- в файле)
	privPem := &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}

	// запись в файл (именно pem, потому что замаркированный объект типа pem)
	pem.Encode(privFile, privPem)

	//ПУБЛИЧНЫЙ КЛЮЧ - исходит от приватного
	pubKey := &privKey.PublicKey

	//файл с публичным ключом
	pubFile, _ := os.Create("public_key.pem")
	defer pubFile.Close()

	// преобразование ключа в стандартизированный вид передачи
	pubBytes := x509.MarshalPKCS1PublicKey(pubKey)

	// маркировка (-BEGIN- -END- в файле)
	pubPem := &pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}

	// запись в файл (именно pem, потому что замаркированный объект типа pem)
	pem.Encode(pubFile, pubPem)

	fmt.Println("Ключи успешно сгенерированы и сохранены в файлы.")
}

// Подпись сообщения с использованием приватного ключа
func signMessage(privKey *rsa.PrivateKey, message string) []byte {
	// Хеширование сообщения с использованием SHA-256, т.к. подписывающая функция требует именно хеш
	hash := sha256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)

	// Подписание хеша сообщения с использованием приватного ключа
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed)
	fmt.Println("Сообщение подписано")
	return signature
}

// Проверка подписи с использованием публичного ключа
func verifySignature(pubKey *rsa.PublicKey, message string, signature []byte) {
	// Хеширование сообщения с использованием SHA-256, т.к. проверяющая функция требует именно хеш
	hash := sha256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)

	// Проверка подписи с использованием публичного ключа
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, signature)

	if err == nil {
		fmt.Println("Подпись подтверждена")
	} else {
		fmt.Println("Подпись недействительна")
	}
}

func main() {

	//создание
	generateKeys()

	message := "Cообщение"

	//извлечение ключей
	privFile, _ := os.ReadFile("private_key.pem")
	block, _ := pem.Decode(privFile)
	if block == nil || block.Type != "PRIVATE KEY" {
		fmt.Errorf("Ключ не является приватным")
	}
	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	pubFile, _ := os.ReadFile("public_key.pem")
	block, _ = pem.Decode(pubFile)
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println("Ключ не является публичным")
	}
	pubKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	// подпись
	signature := signMessage(privKey, message)

	// проверка
	verifySignature(pubKey, message, signature)
}
