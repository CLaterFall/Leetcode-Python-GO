class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        res = []
        length = len(S)
        start = 0
        end = length
        for i in range(length):
            if S[i] == "I":
                res.append(start)
                start += 1
            else:
                res.append(end)
                end -= 1
        res.append(start)
        return res
