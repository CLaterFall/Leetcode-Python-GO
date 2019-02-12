# -*- coding: utf-8 -*-

class Solution:
    def rotatedDigits(self, N: 'int') -> 'int':
        res = 0
        for i in range(1,N+1):
            numStr = str(i)
            if(numStr.count('3') == 0 and numStr.count('4') == 0 and numStr.count('7') == 0):
                if(numStr.count('1') + numStr.count('8') + numStr.count('0') != len(numStr)):
                    res += 1
        return res
