package main

import "fmt"

func main() {
	//nums := []int{1, 1, 1, 1, 2, 3, 4, 4, 4}
	//nums := []int{1, 1, 1, 1}
	//nums := []int{0, 0, 0, 0, 0}
	//nums := []int{1, 1}
	nums := []int{1, 2, 2}
	//nums := []int{1, 2, 3, }
	var l int
	l = removeDuplicates(nums)
	fmt.Println(l)
}

//思路: 1外层循环i, 找i+1 是否相等, 如果相等, 则删除 ,
//		数组集体前移, 再次判断i+1

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for true {
			if len(nums) <= i+1 {
				break
			}
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				break
			}

		}
	}

	return len(nums)
}
