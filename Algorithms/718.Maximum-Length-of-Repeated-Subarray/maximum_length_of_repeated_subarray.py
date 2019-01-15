# -*- coding: utf-8 -*-

class Solution:
    def findLength(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        lenA = len(A)
        lenB = len(B)
        dp = [[0 for j in range(lenB)] for i in range(lenA)]
        res = 0
        for i in range(lenA):
            for j in range(lenB):
                if A[i] == B[j]:
                    if i == 0 or j == 0:
                        dp[i][j] = 1
                    else:
                        dp[i][j] = dp[i - 1][j - 1] + 1
                    res = max(res, dp[i][j])
        return res
