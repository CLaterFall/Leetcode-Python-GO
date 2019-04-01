# -*- coding: utf-8 -*-

class Solution:
    def findCheapestPrice(self, n: 'int', flights: 'List[List[int]]', src: 'int', dst: 'int', K: 'int') -> 'int':
        graph = collections.defaultdict(dict)
        for u, v, w in flights:
            graph[u][v] = w

        res = float('inf')
        que = collections.deque()
        que.append((src, 0))
        step = 0
        while que:
            length = len(que)
            for i in range(length):
                cur, cost = que.popleft()
                if cur == dst:
                    res = min(res, cost)
                for v, w in graph[cur].items():
                    if cost + w > res:
                        continue
                    que.append((v, cost + w))
            if step > K:
                break
            step += 1
        return -1 if res == float('inf') else res
