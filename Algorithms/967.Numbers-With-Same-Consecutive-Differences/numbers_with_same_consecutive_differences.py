class Solution:
    def numsSameConsecDiff(self, N: int, K: int) -> List[int]:
        if N == 1:
            return list(range(10))
        k_list = []
        for i in range(1, 10):
            if i + K < 10:
                k_list.append(i)
            if i - K >= 0:
                k_list.append(i)
    
        for i in range(1, N):
            tmp_res = list()
            for num in k_list:
                pre_ele = num % 10
                if pre_ele + K < 10:
                    tmp_res.append(num * 10 + pre_ele + K)
                if pre_ele - K >= 0:
                    if not (pre_ele - K == 0 and num == 0):
                        tmp_res.append(num * 10 + pre_ele - K)
            k_list = tmp_res         
        return list(set(k_list))   
    
