package main

// 1.	Хэширование данных:
//   •	Разработайте утилиту, которая принимает на вход строку и вычисляет её хэш с использованием алгоритма SHA-256.
//   •	Реализуйте возможность выбора нескольких хэш-функций (например, MD5, SHA-256, SHA-512).
//   •	Включите в утилиту проверку целостности данных: пользователю предлагается ввести строку и её хэш, после чего утилита должна подтвердить или опровергнуть их соответствие.

//Хэширование - это получение уникального идентификатора для набора данных в виде строки единого размера

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

// метод по которому хэшируется строка
func hashing(function string, str string) string {
	var h hash.Hash
	switch function {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		fmt.Println("Неизвестная функция хэширования")
		return ""
	}
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println("Введите строку для хэширования:")
	var str string
	fmt.Scanln(&str)

	fmt.Println("Выберите хэш-функцию:")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA-256")
	fmt.Println("3. SHA-512")
	var choice int
	fmt.Scanln(&choice)

	var function string
	switch choice {
	case 1:
		function = "md5"
	case 2:
		function = "sha256"
	case 3:
		function = "sha512"
	default:
		fmt.Println("Некорректный выбор")
		return
	}

	hashed := hashing(function, str)
	fmt.Printf("Хэш: %s\n", hashed)

	fmt.Print("Введите строку для проверки целостности: ")
	var nStr string
	fmt.Scanln(&nStr)

	fmt.Print("Введите хэш для проверки: ")
	var nFunc string
	fmt.Scanln(&nFunc)

	if hashing(function, nFunc) == nFunc {
		fmt.Println("Целостность данных подтверждена.")
	} else {
		fmt.Println("Целостность данных опровергнута.")
	}
}
