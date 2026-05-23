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
	sum2 := 0
	arr3 := []int{18, 22, 15, 31, 25, 19}
	for i := 0; i < len(arr3); i++ {
		if arr3[i] 
	}
}
