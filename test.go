package main

import (
	"fmt"
)

func main() {
	
	for i := 10; i > 0; i-- {
		fmt.Println(i, "уменьшается")
	}	
	
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[j] * arr[i])
		}
	}

	sum := 0
	for i := 0; 101 > i; i++ {
		sum = sum + i
		fmt.Println(sum)
	}

	arr2 := []int{200, 200, 404, 200, 500, 200, 301, 502}
	for i := 0; i < len(arr2); i++ {
		if arr2[i] >= 400 {
			fmt.Println("error", arr2[i])
		} else {
			fmt.Println("succes", arr2[i])
		}
	}
	arr3 := []int{18, 22, 15, 31, 25, 19}
	max := arr3[0]
	for i := 0; i < len(arr3); i++ {
		if arr3[i] > max {
			max = arr3[i]
		}
	}
	fmt.Println(max)

	arr4 := []string{"Excelent", "Good", "Bad", "Good", "Excelent", "Bad", "Good"}
	/* 
	var good string
	for i := 0; i < len(arr4); i++ {
		if arr4[i] == "Good" {
			good = arr4[i]
		fmt.Println(len(good))
		}
	}
	fmt.Println(len(good))
	*/
	endCount := 0
	for i := 0; i < len(arr4); i++ {
		countNumb := 0
		if arr4[i] == "Good" {
			countNumb++
		}
		endCount = endCount + countNumb
	}
	fmt.Println(endCount)
	
	arr5 := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	for _, v := range arr5 {
		if v % 3 == 0 {
			continue 
		}
		fmt.Println(v)
	}

	for i, v := range arr5 {
		fmt.Println(v)
		fmt.Println(i)
	}
}
