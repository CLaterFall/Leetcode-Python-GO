import collections
class Solution:
    def numPairsDivisibleBy60(self, time: List[int]) -> int:
        res = 0
        cntMap = collections.Counter()
        for ele in time:
            res += cntMap[60 - ele % 60] if ele % 60 else cntMap[0]
            cntMap[ele % 60] += 1
        return res
