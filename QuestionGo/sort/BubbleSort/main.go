package main

import "fmt"

//冒泡排序
func BubbleSort(l []int) []int {
	leng := len(l)
	for indexI, _ := range l {
		for indexL, _ := range l[:leng-indexI-1] {
			if indexL > leng {
				break
			}
			if l[indexL] > l[indexL+1] {
				l[indexL], l[indexL+1] = l[indexL+1], l[indexL]
			}
		}
	}
	return l
}

// 插入排序
func InsertSort(l []int) []int {

	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			for j := i; j > 0; j-- {
				if l[j] < l[j-1] {
					l[j], l[j-1] = l[j-1], l[j]
				}
			}
		}
	}

	return l
}

// 快速排序实现，由于要二分递归，所以要传入每次排序的起始和结束的索引
func quick_sorting(data []int, start, end int) {
	if start < end {
		// 获取基准值
		base := data[start]
		// 临时变量，左右索引
		left := start
		right := end
		// 进入循环
		for left < right {
			// 由于左边的(第0个)被取走当成基准值，所以在左边就留下一个洞，用于存储比基准小的值
			// 所以先从右边找，找到第一个比基准值小的
			for left < right && data[right] >= base {
				right--
			}
			// 找到了比基准值小的，那就把这个值填入左边的洞，做索引要++
			if left < right {
				data[left] = data[right]
				left++
			}
			// 因为上面的操作让右边留了一个洞，开始从左边找比基准值大的
			for left < right && data[left] <= base {
				left++
			}
			// 找到比基准值大的再次把右边洞填上，并且在左边又留了一个洞
			if left < right {
				data[right] = data[left]
				right--
			}
		}

		// 把基准值填入到洞中，这就是本应该他所在的位置
		data[left] = base
		// 继续分两组递归排序，注意此时基准值已经不用参与排序了
		quick_sorting(data, start, left-1)
		quick_sorting(data, left+1, end)
	}
}

func main() {
	list := []int{4, 5, 7, 2, 1, 4, 5, 11}
	//list = BubbleSort(list)
	//list = InsertSort(list)
	fmt.Println(list)
	quick_sorting(list, 0, len(list)-1)
	fmt.Println(list)
}
