class Solution:
    def customSortString(self, S: str, T: str) -> str:
        res = ""
        for ele in S:
            res += ele * T.count(ele)
        
        for ele in T:
            if ele not in S:
                res += ele
        return res
