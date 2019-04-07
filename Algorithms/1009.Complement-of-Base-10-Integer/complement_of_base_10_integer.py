class Solution:
    def bitwiseComplement(self, N: int) -> int:
        if N == 0:
            return 1
        ret = 0
        cnt = 0
        base = 1
        while N:
            cnt  +=1
            #1100110
            #0011001
            #1 0 0 
            if N % 2 == 0:
                ret += base
            base *= 2
            N //= 2
        return ret

