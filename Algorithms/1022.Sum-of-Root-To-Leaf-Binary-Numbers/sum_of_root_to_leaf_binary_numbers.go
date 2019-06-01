/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
    if root == nil {
        return 0
    }
    
    cur = cur * 2 + root.Val
    if root.Left == nil && root.Right == nil {
        return cur
    }
    return dfs(root.Left, cur) + dfs(root.Right, cur)
}