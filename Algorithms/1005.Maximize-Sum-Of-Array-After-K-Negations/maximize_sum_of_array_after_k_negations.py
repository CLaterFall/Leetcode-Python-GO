class Solution:
    def largestSumAfterKNegations(self, A: List[int], K: int) -> int:
        if K == 0:
            return sum(A)
        if len(A) == 0:
            return 0
        
        neg_num = sum([1 if ele < 0 else 0 for ele in A])
        if neg_num > K:
            A.sort()
            for index, num in enumerate(A):
                if index < K:
                    A[index] = -A[index]
                else:
                    break
            return sum(A)
        else:
            flag = (K - neg_num) % 2
            A = list(map(abs, A))
            if flag:
                return sum(A) - min(A)*2
            else:
                return sum(A)
        
