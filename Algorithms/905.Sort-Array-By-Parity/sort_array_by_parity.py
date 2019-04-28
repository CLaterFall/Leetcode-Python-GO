class Solution:
    def sortArrayByParity(self, A: List[int]) -> List[int]:
        res = []
        for ele in A:
            if ele % 2 == 0:
                res.insert(0, ele)
            else:
                res.append(ele)
        return res
