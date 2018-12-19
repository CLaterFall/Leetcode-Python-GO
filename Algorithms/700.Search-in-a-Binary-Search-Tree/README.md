# [700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)

## 题目

Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

For example,

```text
Given the tree:
        4
       / \
      2   7
     / \
    1   3

And the value to search: 2
```

```text
You should return this subtree:

      2
     / \
    1   3
```

In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL.

Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as[], not null.


## 解题思路
BST: 二叉搜索树 -- 若左子树不为空，则左子树上的所有节点的值均小于它的根节点的值，若右子树不为空，则右子树上的所有节点的值均大于它的根节点的值。左右字数也均为二叉搜索树。
故该题只需对根节点，左右子树进行遍历即可。
