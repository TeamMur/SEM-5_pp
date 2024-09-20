добавление числа в массив
append(array, element) возвращает новый массив с добавленным элементом
  nums = append(nums, new_num)

Создание карты сразу с элементами выглядит так
  persons := map[string]int{"Тимур": 20, "Миша": 19}


пробег по ключам мапы происходит так
	for key := range persons {
		age_sum += persons[key]
	}
логику я неочень понимаю, а хотелось бы


деление интов возможно только так и это ужасно
float64(age_sum) / float64(len(persons))