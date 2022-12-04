package statistics_from_a_large_sample

//sampleStats
//  @Description: 该题, 题目不是很清晰, 估计是翻译出问题了 https://leetcode.cn/problems/statistics-from-a-large-sample/
//  @param count
//  @return []float64
//
func sampleStats(count []int) []float64 {
	var minimum, maximum, mean, median, mode float64
	var sum float64
	var m = make(map[int]int)

	for _, val := range count {
		fv := float64(val)
		if minimum != 0 && minimum > fv {
			minimum = fv
		}
		if maximum < fv {
			maximum = fv
		}
		sum = sum + fv
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}
	mean = sum / float64(len(count))
	//排序
	if len(count)%2 == 1 {
		median = float64(count[len(count)-1/2])
	} else {
		median = float64((count[len(count)/2] + count[len(count)/2] - 1) / 2)
	}
	var mo, va int
	for k, v := range m {
		if v > va {
			mo = k
		}
	}
	mode = float64(mo)

	return []float64{minimum, maximum, mean, median, mode}
}
