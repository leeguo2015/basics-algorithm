package checkXMatrix

func checkXMatrix(grid [][]int) bool {
	for i, ints := range grid {
		index, indexReverse := i, len(grid)-i-1
		for key, val := range ints {
			if key == index || key == indexReverse {
				if val == 0 {
					return false
				}
			} else {
				if val != 0 {
					return false
				}
			}
		}
	}
	return true
}
