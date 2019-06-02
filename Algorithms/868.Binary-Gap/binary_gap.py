class Solution:
    def binaryGap(self, N: int) -> int:
        res = 0
        dis = 0
        flag = False
        while N:
            if N % 2 == 1:
                if flag:
                    res = max(res, dis)
                else:
                    flag = True
                dis = 0
            dis += 1
            N = N // 2
        return res
