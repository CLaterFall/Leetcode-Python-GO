from math import log
class Solution:
    def powerfulIntegers(self, x: int, y: int, bound: int) -> List[int]:
        res = set()
        if bound <= 1:
            return []
        lgx = int(log(bound, max(x, 2))) + 1
        lgy = int(log(bound, max(y, 2))) + 1
        for i in range(lgx):
            for j in range(lgy):
                ssum = x ** i + y ** j
                if ssum > bound:
                    break
                res.add(ssum)
        return list(res)

