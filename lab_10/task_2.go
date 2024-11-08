package main

// 2.	Симметричное шифрование:
//   •	Реализуйте программу, шифрующую переданные данные с помощью алгоритма AES.
//   •	Пользователь должен указать строку и секретный ключ.
//   •	Программа должна зашифровать строку и предоставить возможность расшифровать её при вводе того же ключа.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

// Кодирует строку по алгоритму AES
func encrypt(s string, k string) string {

	// Ключ должен быть в типе byte[] и кратен 16, 24, или 32 (прописано в методах aes)
	// Возьмем хеш md5, который кратен 16
	key := []byte(k)
	md5 := md5.Sum(key) // md5.Sum возвращает массив, а не срез
	key = md5[:]        // Создаем срез
	block, _ := aes.NewCipher(key)

	// Строка тоже должна быть в типе byte и быть кратна тому чему кратен ключ
	str := []byte(s)

	// Дозаполнение (по сути наполнение строки пробелами, а вернее пустыми символами)
	diff := aes.BlockSize - (len(str) % aes.BlockSize)
	space := make([]byte, diff)
	str = append(str, space...)

	// Вектор для шифрования. Те же критерии как к ключу и строке
	// ПОНЯТЬ: для чего это?
	iv := make([]byte, aes.BlockSize)
	// rand.Read(iv) - везде его вставляют, но даже без рандома работает

	// ПОНЯТЬ: *шифрование*
	mode := cipher.NewCBCEncrypter(block, iv)

	// ПОНЯТЬ: *шифрование*
	ciphertext := make([]byte, len(str))
	mode.CryptBlocks(ciphertext, str)

	// Возвращаем строку в представлении base64, предварительно объединив входную строку и закодированный текст
	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result)
}

// Декодирует строку, закодированную по алгоритму AES
func decrypt(s string, k string) string {
	// преобразуем base64-строку в срез байт
	str, _ := base64.StdEncoding.DecodeString(s)

	// Ключ должен быть в типе byte[] и кратен 16, 24, или 32 (прописано в методах aes)
	// Возьмем хеш md5, который кратен 16
	key := []byte(k)
	md5 := md5.Sum(key) // md5.Sum возвращает массив, а не срез
	key = md5[:]        // Создаем срез
	block, _ := aes.NewCipher(key)

	// Извлекаем зашифрованный текст и входной вектор для шифрования
	ciphertext := str[aes.BlockSize:]
	iv := str[:aes.BlockSize]

	// Создаем срез в который будет помещена наша декодированная строка
	text := make([]byte, len(ciphertext))

	// ПОНЯТЬ: Дешифрование
	// и всё же зачем iv
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(text, ciphertext)

	// Получаем исходную строку убирая пустые символы, добавленные при шифровании
	result := string(text)
	var c byte = 4 //символ пустоты
	return strings.ReplaceAll(result, string(c), "")
}

func main() {
	enc := encrypt("Привет Тимур", "ключ")
	fmt.Println(enc)
	fmt.Println(decrypt(enc, "ключ"))
}
