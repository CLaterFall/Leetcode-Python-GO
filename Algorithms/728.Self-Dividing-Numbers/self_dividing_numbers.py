class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        res = []
        for ele in range(left, right + 1):
            temp = ele
            flag = True
            while temp:
                remain = temp % 10
                temp = temp // 10
                if remain == 0 or ele % remain != 0:
                    flag = False
                    break
            if flag:
                res.append(ele)
        return res
