"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        stack = [root]
        res = []
        while stack:
            tmp = stack.pop(0)
            if tmp:
                res.append(tmp.val)
                if tmp.children:
                    stack = tmp.children + stack
        return res
