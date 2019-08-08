class Solution:
    def numJewelsInStones(self, J: str, S: str) -> int:
        return sum(S.count(ele) for ele in J)
