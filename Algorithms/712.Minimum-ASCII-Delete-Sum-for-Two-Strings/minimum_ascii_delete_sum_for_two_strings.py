# -*- coding: utf-8 -*-

class Solution:
    def minimumDeleteSum(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: int
        """
        len1, len2 = len(s1), len(s2)
        dp = [[0] * (len2 + 1) for _ in range (len1 + 1)]
        for i in range(len1):
            dp[i + 1][0] = dp[i][0] + ord(s1[i])
        for j in range(len2):
            dp[0][j + 1] = dp[0][j] + ord(s2[j])
        for i in range(1, len1 + 1):
            for j in range(1, len2 + 1):
                if s1[i - 1] == s2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = min(dp[i - 1][j] + ord(s1[i - 1]), dp[i][j - 1] + ord(s2[j - 1]))
        return dp[-1][-1]
        
