package main

import "fmt"

func main() {
	name := Greet("alex")
	course := EuroToRub(49.3)
	folder_course := &course
	age := IsAdult(19)
	a, b := MinMax(3, 12)
	d1, d2 := Divide(13, 4)
	calc := Calc(folder_course)
	arr := []int{1,2,3,4,5,6,7,8,9,10,12,13,231,321}
	sumArr := SumArray(arr)
	test1, test2 := SafeDivide(13, 0)
	arrT := []string{"hello", "bye", "run", "fast"}
	target := "run"
	arr1 := Contains(arrT, target)
	
	fmt.Println("Ваше имя: ", name)
	fmt.Println("Ваш баланс: ", course)
	fmt.Println("Ваш возраст", age)
	fmt.Println("Большее число: ",a, b)
	fmt.Println("Деление чисел: ", d1, d2)
	fmt.Println("Это калк: ", calc)
	fmt.Println("Отсортированное", sumArr)
	fmt.Println("Тест на ошибку", test1, test2)
	fmt.Println("Тестинг", arr1)
}

func Greet(name string) string {
	return name
}

func EuroToRub(euros float64) float64 {
	price := 90.5
	sum := euros * price
	return sum
}

func IsAdult(age int) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}

func MinMax(a, b int) (int, int) {
	if a > b {
		return a, b
	} else if a < b {
		return b, a
	} else {
		return b, a
	}
}

func Divide(a, b int) (int, int) {
	res1 := a/b
	res2 := a%b
	return res1, res2
}

func Calc(a *float64) (courses float64) {
	courses = *a
	courses = courses - 123
	return 
}

func SumArray(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
	sum = sum + numbers[i]
	}
	return sum
}

func SafeDivide(a, b float64) (float64, error) {
	sum := a / b
	if b == 0 {
		fmt.Errorf("деление на 0")
	}
	return sum, nil
}

func Contains(arr []string, target string) bool {
	test := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			test = true
		}
	}
	return test
}
