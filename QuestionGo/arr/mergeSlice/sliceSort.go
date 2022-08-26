package mergeSlice

import "fmt"

func SliceSort(s1 []int, s2 []int) []int {
	newSlice := append(s1, s2...)
	fmt.Println(newSlice)
	return newSlice
}
