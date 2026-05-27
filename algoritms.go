package main

import (
	"fmt"
	"sort"
)

func main() {
	one := "top"
	two := "pot"
	fmt.Println(IsAnagram(one, two))

	arr := []int{1, 3, 5, 7, 9, 11, 15}
	target := 16
	fmt.Println(TargetNumb(arr, target))
}

func IsAnagram(one string, two string) bool {
	if len(one) != len(two) {
		return false
	}

	arr1 := []rune(one)
	arr2 := []rune(two)

	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	return string(arr1) == string(arr2)
}

func TargetNumb(arr []int, target int) (int, int) {
	l := 0
	r := len(arr) - 1
	for l < r {
		sum := arr[l] + arr[r]
		if sum == target {
			return l, r
		} else if sum > target {
			r--
		} else if sum < target {
			l++
		}
	}
	return -1, -1
}

