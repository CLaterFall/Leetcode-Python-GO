class Solution:
    def peakIndexInMountainArray(self, A: List[int]) -> int:
        length = len(A)
        low = 0
        high = length - 1
        while low < high:
            mid = (low + high) // 2
            if A[mid] > A[mid - 1] and A[mid] > A[mid + 1]:
                return mid
            elif A[mid - 1] < A[mid] < A[mid + 1]:
                low = mid
            elif A[mid-1] > A[mid] > A[mid+1]:
                high = mid
