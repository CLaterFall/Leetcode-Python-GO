# -*- coding: utf-8 -*-


class Solution:
    def isMonotonic(self, A: List[int]) -> bool:
        flag = None
        for i in range(len(A) - 1):
            diff = A[i] - A[i+1]
            if diff == 0:
                continue
            diff = diff < 0
            if flag == None:
                flag = diff
            if diff != flag:
                return False
        return True
