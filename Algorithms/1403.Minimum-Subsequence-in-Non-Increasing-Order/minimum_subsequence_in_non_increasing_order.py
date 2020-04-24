class Solution:
    def minSubsequence(self, nums: List[int]) -> List[int]:
        total = sum(nums)
        nums.sort(reverse=True)
        res = []
        cur = 0
        for ele in nums:
            cur += ele
            if cur > total - cur:
                res.append(ele)
                break
            res.append(ele)
        return res
