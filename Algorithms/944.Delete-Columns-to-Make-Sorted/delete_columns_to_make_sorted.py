# -*- coding: utf-8 -*-

class Solution:
    def minDeletionSize(self, A: List[str]) -> int:
        res = 0
        length = len(A[0])
        for i in range(length):
            col = [ele[i] for ele in A]
            if col != sorted(col):
                res +=1
        return res
