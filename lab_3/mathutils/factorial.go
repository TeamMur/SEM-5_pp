package mathutils

func GetFactorial(num int) int {
	if num <= 1 {
		return 1
	}
	return GetFactorial(num-1) * num
}
