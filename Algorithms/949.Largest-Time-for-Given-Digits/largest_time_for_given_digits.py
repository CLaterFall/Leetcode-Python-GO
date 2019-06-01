import itertools
class Solution:
    def largestTimeFromDigits(self, A: List[int]) -> str:
        permutations = [p for p in itertools.permutations(A) if p <= (2, 3, 5, 9)]
        permutations.sort()
        length = len(permutations)
        res = ""
        while length:
            ele = permutations[length - 1]
            if ele[2:] <= (5, 9):
                res = '%d%d:%d%d' % ele
                return res
            length -= 1
        return res
