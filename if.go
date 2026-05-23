package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,99,7,7,5,234,32,46,545,23,434,5,3424,134,35,23,43}
	for i := 0; len(arr) > i; i++ {
		if arr[i] % 2 == 0 {
			fmt.Println(arr[i], "чётнок число")
		} else {
			fmt.Println(arr[i], "не чётное число")
		}
	}

	age := []int{4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30}
	for i := 0; len(age) > i; i++ {
		if age[i] >= 18 {
			fmt.Println("Вам больше 18")
		} else {
			fmt.Println("Вам меньше 18")
		}
	}

	first_password := bufio.NewScanner(os.Stdin)
	second_password := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите пароль, повторите его")
	first_password.Scan()
	second_password.Scan()
	if first_password.Text() == second_password.Text() {
		fmt.Println("ваш пароль совпадает")
	} else {
		fmt.Println("ваш пароль не совпадает")
	}
	
	IsAdmin := false
	IsRoot := true

	if IsAdmin || IsRoot {
		fmt.Println("Система доступна для настройки")
	} else {
		fmt.Println("Система недоступна для настройки")
	}
	
	fmt.Println("Введите порт")
	port := bufio.NewScanner(os.Stdin)
	port.Scan()
	data := port.Text()
	data = strings.TrimSpace(data)
	port_check, err := strconv.Atoi(data)
	if err != nil {
		fmt.Println("Вводи число гений")
		return
	}
	if port_check > 1024 && port_check < 65535 {
		fmt.Println("порт доступен")
	} else {
		fmt.Println("порт недоступен")
	}
	
	fmt.Println("Введите цвет светофора на английском")
	colorChoise := bufio.NewScanner(os.Stdin)
	colorChoise.Scan()
	colorChoiseData := strings.TrimSpace(colorChoise.Text())
	if colorChoiseData == "green" {
		fmt.Println("Иди")
	} else if colorChoiseData == "yellow" {
		fmt.Println("Жди")
	} else if colorChoiseData == "red" {
		fmt.Println("Стой")
	} else {
		fmt.Println("Светофор сломался")
	}
	
	fmt.Println("Проверка скорости введите вашу скорость")
	new_speed := bufio.NewScanner(os.Stdin)
	new_speed.Scan()
	new_speed_data := new_speed.Text()
	new_speed_data = strings.TrimSpace(new_speed_data)
	new_speed_check, err := strconv.Atoi(new_speed_data)
	if err != nil {
		fmt.Println("Ошибка вы не ввели число")
		return
	}
	if x := new_speed_check; x > 90 {
		fmt.Println("Вы превысили скорость у вас штраф")
	} else {
		fmt.Println("Вы не превысили скорость у вас нету штраф")
	}
	fmt.Println("Создайте свой счёт")
	balance := bufio.NewScanner(os.Stdin)
	balance.Scan()
	balance_data := balance.Text()
	balance_data = strings.TrimSpace(balance_data)
	balance_check, err := strconv.Atoi(balance_data)
	if err != nil {
		fmt.Println("Введити числа с запятой")
		return
	}
	balance_result := float64(balance_check)
	const cost float64 = 17.5
	fmt.Println("Ваш баланс", balance_result, "Стоимость звонка", cost)
	if balance_result > 0 && cost < balance_result {
		balance_result = balance_result - cost
		fmt.Println("Вы сделали звонок")
		fmt.Println("Ваш баланс", balance_result)
	} else {
		fmt.Println("На вашем счету не хватает средств")
	}

	status := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите ваш статус от 200 до 500")
	status.Scan()
	status_data := status.Text()
	status_data = strings.TrimSpace(status_data)
	status_check, err := strconv.Atoi(status_data)
	if err != nil {
		fmt.Println("Статус: Ошибка!!!")
	}

	status_result := int(status_check)
	if status_result > 200 && status_result < 300 {
		fmt.Println("Статус успешен", status_result)
	} else if status_result > 300 && status_result < 400 {
		fmt.Println("Статус ошибка клиента", status_result)
	} else if status_result > 400 && status_result < 500 {
		fmt.Println("Статус ошибка сервера", status_result)
	} else {
		fmt.Println("Неизвестная ошибка")
	}




}
