# -*- coding:utf-8 _*-
"""
@author:lee 
@file: optimal_shares.py 
@time: 2021/05/11
@contact: leeguo2015@163.com
@site:  
@software: PyCharm 

"""
"""
# 贪心算法

给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

"""


class Solution:

    def maxProfit(self, nums: list, principal: int) -> int:

        for i in nums:
            to_day = nums.index(i)
            remain_day = nums[to_day:]
            profit = 0

            for j in remain_day:
                if j - i and j-i > profit:
                    profit = j - i

        return len(nums)


if __name__ == "__main__":
    data = [1, 2, 2, 3]
    result = Solution().maxProfit(data)
    print(result, data)
