package main

import (
	"fmt"
	"sort"
)

func main() {
	one := "top"
	two := "pot"
	fmt.Println(IsAnagram(one, two))

	numb_one := []int{1,1,2,4,2,1,3,2,3,2,0,1,0,3,2,3,1,0}
	fmt.Println(CountNumb(numb_one))
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

func CountNumb(one []int{}) []int{} {
	sort.Slice(one)
	for i:=0, len(one) > i; i++ {
		

}

