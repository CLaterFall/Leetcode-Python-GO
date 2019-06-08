class Solution:
    def convertToBase7(self, num: int) -> str:
        res = ""
        sig = num//abs(num) if num!=0 else 0
        num = abs(num)
        while num >= 7:
            cur_num = num % 7
            num = num // 7
            res = str(cur_num)+res
        res = str(num) + res
        if sig >= 0:
            return res
        else:
            return "-" + res
