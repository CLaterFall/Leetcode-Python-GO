class Solution:
    def calPoints(self, ops: List[str]) -> int:
        stack = []
        for ele in ops:
            if ele == "C":
                stack.pop()
            elif ele == "D":
                stack.append(stack[-1] * 2)
            elif ele == "+":
                stack.append(stack[-1] + stack[-2])
            else:
                num = int(ele)
                stack.append(num)
        return sum(stack)
