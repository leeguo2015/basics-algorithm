package insertionSort

/*
插入排序
*/

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for inner := 0; inner < len(arr[:i]); inner++ {
			if arr[i] < arr[inner] {
				arr = append(arr[:inner-1], arr[i])
				arr = append(arr, arr[inner:]...)
			}
		}

	}
	return arr
}
