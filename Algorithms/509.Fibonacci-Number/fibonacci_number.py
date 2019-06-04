class Solution:
    def fib(self, N: int) -> int:
        f = [0, 1]
        pos = 2
        while pos <= N:
            f.append(f[-1] + f[-2])
            pos += 1
        return f[N]
