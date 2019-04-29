# [908. Smallest Range I](https://leetcode.com/problems/smallest-range-i/)

## 题目
Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i].

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.


Example 1:
```text
Input: A = [1], K = 0
Output: 0
Explanation: B = [1]
```
Example 2:
```text
Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]
```
Example 3:
```text
Input: A = [1,3,6], K = 3
Output: 0
Explanation: B = [3,3,3] or B = [4,4,4]
```

Note:

1. 1 <= A.length <= 10000
2. 0 <= A[i] <= 10000
3. 0 <= K <= 10000


## 解题思路
要求的是新的数组最大值和最小值之差，所以，我们应该把所有的数字向中间靠拢。即A中的最小值+K, 最大值-K，判断这样操作之后，能否有重叠，如果能重叠所求的结果就是0；如果不能重叠，所求的结果就是两者的差值。
