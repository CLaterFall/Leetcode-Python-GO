class Solution:
    def sortedSquares(self, A: List[int]) -> List[int]:
        return sorted([ele**2 for ele in A])
