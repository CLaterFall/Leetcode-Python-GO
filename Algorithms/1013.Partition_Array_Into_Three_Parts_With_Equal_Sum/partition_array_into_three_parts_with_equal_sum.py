class Solution:
    def canThreePartsEqualSum(self, A: List[int]) -> bool:
        total = sum(A)
        if total % 3 != 0:
            return False
        
        equal = total // 3
        part = 0
        count = 0
        i = 0
        j = 0
        for idx in range(0, len(A)):
            part += A[idx]
            if part == equal:
                part = 0
                count += 1
                if count == 1 :
                    i = idx
                else:
                    j = idx + 1
                
        if count == 3 and i + 1 < j:
            return True
        return False
            
