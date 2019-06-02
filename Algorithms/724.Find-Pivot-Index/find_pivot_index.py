class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        left, total = 0, sum(nums)
        for index, num in enumerate(nums):
            if left * 2 + num == total:
                return index
            left += num
        return -1

