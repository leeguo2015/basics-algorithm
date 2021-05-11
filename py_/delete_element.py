# -*- coding:utf-8 _*-
"""
@author:lee 
@file: delete_element.py 
@time: 2021/05/10
@contact: leeguo2015@163.com
@site:  
@software: PyCharm 

"""

"""
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

"""


class Solution:

    def removeDuplicates(self, nums: list) -> int:

        # while True:
        #     for i in nums:
        #         if nums.index(i) == len(nums) - 1:
        #             continue
        #         if i == nums[nums.index(i) + 1]:
        #             del nums[nums.index(i) + 1]
        #             break
        #     else:
        #         return len(nums)

        for i in nums:
            if nums.index(i) == len(nums) - 1:
                continue
            for j in nums[nums.index(i):]:
                if i == j:
                    del nums[nums.index(i) + 1]

        else:
            return len(nums)


if __name__ == "__main__":
    data = [1, 2, 2, 3]
    # data = [1, 1, 1, 1]
    result = Solution().removeDuplicates(data)
    print(result, data)
