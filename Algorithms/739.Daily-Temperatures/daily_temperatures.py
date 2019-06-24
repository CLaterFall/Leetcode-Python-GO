class Solution:
    def dailyTemperatures(self, T: List[int]) -> List[int]:
        temp = []        
        length = len(T)
        res = [0] * length
        
        temp.append(0)
        for i in range(length):
            while temp and T[i] > T[temp[-1]]:
                idx = temp.pop()
                res[idx] = i - idx
            temp.append(i)
        return res
