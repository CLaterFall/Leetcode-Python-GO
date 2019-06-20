class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        five_cnt = 0
        ten_cnt = 0
        for ele in bills:
            if ele == 5:
                five_cnt += 1
            elif ele == 10:
                if five_cnt < 1:
                    return False
                five_cnt -= 1
                ten_cnt += 1
            else:
                if ten_cnt > 0:
                    ten_cnt -= 1
                    five_cnt -= 1
                else:
                    five_cnt -= 3
            if five_cnt < 0 or ten_cnt < 0:
                return False
        return True
