package main

import "fmt"

func main() {
	//nums := []int{1, 1, 1, 1, 2, 3, 4, 4, 4}
	nums := []int{1, 1, 1, 1}
	//nums := []int{1, 1, 1, 1}
	//nums := []int{1, 2, 3, }
	var l int
	l = removeDuplicates(nums)
	fmt.Println(l)
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				break
			}
		}
	}
	return len(nums)
}

//func removeDuplicates(nums []int) int {
//	fmt.Println(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		for j := 0; j < len(nums)-i; j++ {
//			fmt.Printf("第一位i下标%v, 值为%v, 子数组j 下标%v, 值为%v \n", i, nums[i], j, nums[i+j])
//			if nums[i] == nums[i+1] {
//				fmt.Printf("删除的元素%v ,删除元素下标为%v, 前的数组为: %v\n", nums[i+j], i+j, nums)
//				nums = append(nums[:i], nums[i+1:]...)
//				fmt.Printf("后的数组为: %v\n", nums)
//			} else {
//				fmt.Printf("不删除的元素%v ,删除元素i下标为%v,  前的数组为: %v 不删除的j元素%v ,删除元素j下标为%v,\n", nums[i], i, nums, nums[i+j], i+j)
//				break
//			}
//		}
//
//	}
//	fmt.Println(nums)
//
//	return len(nums)
//}
