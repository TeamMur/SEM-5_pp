package main

import (
	"fmt"
	"math"
)

func main() {
	//1. Создать структуру Person с полями name и age. Реализовать метод для вывода информации о человеке.
	fmt.Println("1. Вывод информации о человеке:")
	var my_person Person = Person{name: "Тимур", age: 20}
	print_person_info(my_person)

	fmt.Println()
	//2. Реализовать метод birthday для структуры Person, который увеличивает возраст на 1 год.
	fmt.Println("2. Результат использования birthday():")
	birthday(&my_person)
	fmt.Println(my_person.age)

	fmt.Println()
	//3. Создать структуру Circle с полем radius и метод для вычисления площади круга.
	fmt.Println("3. Площадь круга:")
	fmt.Print("Введите радиус круга: ")
	var radius float64
	fmt.Scan(&radius)
	var circle Circle = Circle{radius: radius}
	third_result := get_circle_area(circle)
	fmt.Printf("Результат: %0.2f", third_result)

	fmt.Println()
	//4. Создать интерфейс Shape с методом Area(). Реализовать этот интерфейс для структур Rectangle и Circle.

	//5. Реализовать функцию, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта.

	//6. Создать интерфейс Stringer и реализовать его для структуры Book, которая хранит информацию о книге.

	fmt.Println("\nКонец выполнения программы")
}

// 1.
type Person struct {
	name string
	age  int
}

func print_person_info(person Person) {
	fmt.Println("Имя: ", person.name, " Возраст: ", person.age)
}

// 2.
func birthday(person *Person) {
	person.age += 1
}

// 3.
type Circle struct {
	radius float64
}

func get_circle_area(circle Circle) float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
