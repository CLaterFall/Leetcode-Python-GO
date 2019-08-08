class Solution:
    def rotateString(self, A: str, B: str) -> bool:
        return (B in A+A and len(B) == len(A))
