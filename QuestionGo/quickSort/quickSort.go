package quickSort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	indexP := arr[0]
	leftList := []int{}  // make([]int, 1)
	rightList := []int{} // make([]int, 1)
	sameList := []int{}
	for i := 0; i < len(arr); i++ {
		if indexP > arr[i] {
			leftList = append(leftList, arr[i])
		} else if indexP < arr[i] {
			rightList = append(rightList, arr[i])
		} else {
			sameList = append(sameList, arr[i])
		}
	}
	leftList = append(QuickSort(leftList), sameList...)
	rightList = append(leftList, QuickSort(rightList)...)
	return rightList
}
