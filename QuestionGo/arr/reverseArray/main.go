package main

func rotate(nums []int, k int) {
	l := len(nums)

	for i := 0; i < l; i++ {

		if i > l-k {
			break
		}

		nowIndex := i
		newIndex := i + k
		if newIndex >= l {
			newIndex = newIndex - l
		}
		//fmt.Println(i, l, k, nowIndex, newIndex)
		t := nums[nowIndex]
		nums[nowIndex] = nums[newIndex]
		nums[newIndex] = t
	}
	//return nums
}

func main() {
	//fmt.Println("a")
	rotate([]int{
		1, 2, 3, 4, 5, 6, 7,
	}, 3)
}
