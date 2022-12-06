package closestCost

import "math"

/*
https://leetcode.cn/problems/closest-dessert-cost/description/
1774. 最接近目标价格的甜点成本
你打算做甜点，现在需要购买配料。目前共有 n 种冰激凌基料和 m 种配料可供选购。而制作甜点需要遵循以下几条规则：

必须选择 一种 冰激凌基料。
可以添加 一种或多种 配料，也可以不添加任何配料。
每种类型的配料 最多两份 。
给你以下三个输入：

baseCosts ，一个长度为 n 的整数数组，其中每个 baseCosts[i] 表示第 i 种冰激凌基料的价格。
toppingCosts，一个长度为 m 的整数数组，其中每个 toppingCosts[i] 表示 一份 第 i 种冰激凌配料的价格。
target ，一个整数，表示你制作甜点的目标价格。
你希望自己做的甜点总成本尽可能接近目标价格 target 。

返回最接近 target 的甜点成本。如果有多种方案，返回 成本相对较低 的一种。
*/

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans := math.MaxInt
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	var dfs func(idx int, sum int, target int, toppingCosts []int)
	dfs = func(idx int, sum int, target int, toppingCosts []int) {
		dis := abs(target - sum)
		dis2 := abs(target - ans)
		if dis < dis2 || (dis == dis2 && sum < ans) {
			ans = sum
		}
		if sum > target {
			return
		}
		for i := idx; i < len(toppingCosts); i++ {
			dfs(i+1, sum+toppingCosts[i], target, toppingCosts)
			dfs(i+1, sum+2*toppingCosts[i], target, toppingCosts)
		}
	}

	for _, b := range baseCosts {
		dfs(0, b, target, toppingCosts)
	}
	return ans
}

//func abs(x int) int {
//	if x > 0 {
//		return x
//	}
//	return -x
//}

//func closestCost(baseCosts []int, toppingCosts []int, target int) int {
//	var bastVal int
//	for _, baseCost := range baseCosts {
//		if bastVal == 0 {
//			bastVal = baseCost
//		}
//		for baseCount := 1; true; baseCount++ {
//			if theVal := baseCount * baseCost; target > theVal && theVal >= bastVal {
//				bastVal = theVal
//			}
//			if theVal := baseCount * baseCost; theVal > baseCost {
//				break
//			}
//			for _, topping := range toppingCosts {
//				for  {
//
//				}
//			}
//		}
//
//	}
//	return bastVal
//}
//
//func closestCost()  {
//
//}
