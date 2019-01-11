# -*- coding: utf-8 -*-

class Solution:
    def numSubarrayProductLessThanK(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        length = len(nums)
        left, right = 0, 0
        res = 0
        prod = 1
        while right < length:
            prod *= nums[right]
            while left <= right and prod >= k:
                prod /= nums[left]
                left += 1
            res += right - left + 1
            right += 1
        return res
