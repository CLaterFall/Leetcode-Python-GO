class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        q = []
        for i, r in enumerate(mat):
            num = sum(r)
            heapq.heappush(q, (-num, -i))
            if len(q) > k:
                heapq.heappop(q)
        return [-t[1] for t in heapq.nlargest(k, q)]
