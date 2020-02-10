class Solution:
    def countPrimeSetBits(self, L: int, R: int) -> int:
        res = 0
        for i in range(L, R+1):
            count = 0
            for j in bin(i):
                if j == "1":
                    count += 1
            if count in {2, 3, 5, 7, 11, 13, 17, 19}:
                res += 1
        return res