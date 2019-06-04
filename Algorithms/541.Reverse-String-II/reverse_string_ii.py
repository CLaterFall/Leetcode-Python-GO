class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        p = 0
        length = len(s)
        while p*k < length:
            s = s[:p*k] + s[p*k: p*k+k][::-1] + s[p*k+k:]
            p += 2
        return s
