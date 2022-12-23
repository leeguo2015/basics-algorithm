package finalValueAfterOperations

/*
https://leetcode.cn/problems/final-value-of-variable-after-performing-operations/description/
存在一种仅支持 4 种操作和 1 个变量 X 的编程语言：

++X 和 X++ 使变量 X 的值 加 1
--X 和 X-- 使变量 X 的值 减 1
最初，X 的值是 0

给你一个字符串数组 operations ，这是由操作组成的一个列表，返回执行所有操作后， X 的 最终值 。
*/
func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, operation := range operations {
		switch operation {
		case "++X", "X++":
			result++
		case "--X", "X--":
			result--
		}
	}
	return result
}
