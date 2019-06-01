# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumRootToLeaf(self, root: TreeNode) -> int:
        if not root:
            return 0
        self.res = 0
        self.dfs(root, root.val)
        return self.res
    
    
    def dfs(self, root, curSum):
        if not root.left and not root.right:
            self.res = self.res + curSum
            return
        if root.left:
            self.dfs(root.left, curSum * 2 + root.left.val)
        if root.right:
            self.dfs(root.right, curSum * 2 + root.right.val)
