class Solution:
    def repeatedNTimes(self, A: List[int]) -> int:
        ele = set()
        for i in A:
            if i in ele:
                return i
            ele.add(i)
        return -1
        
