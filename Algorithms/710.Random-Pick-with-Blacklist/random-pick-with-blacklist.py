# -*- coding: utf-8 -*-
import random


class Solution:

    def __init__(self, N, blacklist):
        """
        :type N: int
        :type blacklist: List[int]
        """
        self.length = N - 1
        self.dic = {}
        blacklist.sort(reverse=True)
        for ele in blacklist:
            self.dic[ele] = self.dic.get(self.length, self.length)
            self.length -= 1
        

    def pick(self):
        """
        :rtype: int
        """
        num = random.randint(0, self.length)
        return self.dic.get(num, num)
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(N, blacklist)
# param_1 = obj.pick()
