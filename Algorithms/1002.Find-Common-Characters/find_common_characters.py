from collections import Counter
class Solution:
    def commonChars(self, A: List[str]) -> List[str]:
        if len(A) == 0:
            return []
        temp = Counter(A[0])
        for i in range(1, len(A)):
            temp &= Counter(A[i])
        return list(temp.elements())
