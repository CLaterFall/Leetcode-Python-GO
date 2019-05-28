# [1054. Distant Barcodes](https://leetcode.com/problems/distant-barcodes/)

## 题目
In a warehouse, there is a row of barcodes, where the i-th barcode is barcodes[i].

Rearrange the barcodes so that no two adjacent barcodes are equal.  You may return any answer, and it is guaranteed an answer exists.

 

Example 1:
```text
Input: [1,1,1,2,2,2]
Output: [2,1,2,1,2,1]
```
Example 2:
```text
Input: [1,1,1,1,2,2,3,3]
Output: [1,3,1,3,2,1,2,1]
```

Note:

1. 1 <= barcodes.length <= 10000
2. 1 <= barcodes[i] <= 10000


## 解题思路
首先根据元素出现次数多少进行排序，然后填充偶数位，再填充基数位。