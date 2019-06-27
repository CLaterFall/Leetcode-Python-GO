# [796. Rotate String](https://leetcode.com/problems/rotate-string/)

## 题目
We are given two strings, A and B.

A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example, if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if A can become B after some number of shifts on A.

Example 1:
```text
Input: A = 'abcde', B = 'cdeab'
Output: true
```

Example 2:
```text
Input: A = 'abcde', B = 'abced'
Output: false
```

Note:

- A and B will have length at most 100.

## 解题思路
A和B的长度必须一样，B必定在AA之中。