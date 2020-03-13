class Solution:
    def maxFreq(self, s: str, maxLetters: int, minSize: int, maxSize: int) -> int:
        dic = {}
        length = len(s)
        if length < minSize:
            return 0
      
        res = 0
        for i in range(minSize, length+1):
            for j in range(minSize, maxSize+1):
                if i - j >= 0:
                    tmp = s[i-j: i]
                    if len(set(tmp)) <= maxLetters:
                        if tmp in dic:
                            dic[tmp] += 1
                        else:
                            dic[tmp] = 1
                        res = max(res, dic[tmp])
                else:
                    break
        
      
        return res
            
            

