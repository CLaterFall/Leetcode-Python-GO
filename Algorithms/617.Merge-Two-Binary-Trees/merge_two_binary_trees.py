# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        def mergeTree( t11: TreeNode, t22: TreeNode) -> TreeNode:
            if t11 and t22:
                t11.val+=t22.val
            elif (not t11) and t22:
                t11 = t22
                t22 = None
                return t11
            else:
                return t11    
            t11.left = mergeTree(t11.left,t22.left)
            t11.right = mergeTree(t11.right,t22.right)
            return t11
        MT = mergeTree(t1,t2)
        return MT
