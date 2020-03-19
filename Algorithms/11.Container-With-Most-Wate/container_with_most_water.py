class Solution:
    def maxArea(self, height: List[int]) -> int:
        head = 0
        tail = len(height) - 1
        ret = 0
        while head != tail:
            if height[head] <= height[tail]:
                ret = max(ret, height[head]*(tail-head))
                head += 1
            else:
                ret = max(ret, height[tail]*(tail-head))
                tail -= 1
        return ret
