import collections
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        dic = {"b": 1, "a": 1, "l": 2, "o": 2, "n": 1}
        counter = collections.Counter(text)
        res = None
        foot = dict()
        for k, v in counter.items():
            if k in dic:
                foot[k] = 1
                cnt = v // dic[k]
                if res == None:
                    res = cnt
                else:
                    res = min(res, cnt)
        for k in dic.keys():
            if k not in foot:
                return 0
        if res == None:
            return 0
        return res
