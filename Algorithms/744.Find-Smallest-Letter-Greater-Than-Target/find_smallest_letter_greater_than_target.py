class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        length = len(letters)
        start = 0
        end = length - 1
        while start <= end:
            mid = (start + end) // 2
            if letters[mid] > target:
                end = mid - 1
            else:
                start = mid + 1
        return  letters[0] if start >= length else letters[start]
