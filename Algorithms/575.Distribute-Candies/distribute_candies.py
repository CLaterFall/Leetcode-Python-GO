import math
class Solution:
    def distributeCandies(self, candies: List[int]) -> int:
        sorts = len(set(candies))
        cnt = len(candies)
        if sorts < cnt / 2:
            return sorts
        else:
            return math.ceil(cnt / 2)
