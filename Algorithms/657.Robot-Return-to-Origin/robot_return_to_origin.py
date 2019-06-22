class Solution:
    def judgeCircle(self, moves: str) -> bool:
        res = [0, 0]
        for ele in moves:
            if ele == "U":
                res[1] += 1
            elif ele == "D":
                res[1] -= 1
            elif ele == "L":
                res[0] -= 1
            else:
                res[0] += 1
        return res[0] == 0 and res[1] == 0
                
