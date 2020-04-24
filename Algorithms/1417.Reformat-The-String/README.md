# [1417. Reformat The String](https://leetcode.com/problems/reformat-the-string/)


## 题目
Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.


Example 1:
```
Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
```
Example 2:
```
Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.
```
Example 3:
```
Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.
```
Example 4:
```
Input: s = "covid2019"
Output: "c2o0v1i9d"
```
Example 5:
```
Input: s = "ab123"
Output: "1a2b3"
```

Constraints:

- 1 <= s.length <= 500
- s consists of only lowercase English letters and/or digits.


## 解题思路
python 解法用了一个比较小聪明的做法，建了个数组为输入字符串长度+1,数字字符从1位开始添加，字母字符从2位开始添加，0位预留，如果字母字符比数字字符多，则回写到0位,每次都记录的位置，大于length+1则返回空字符串，否则返回数组拼装的字符串。
go 解法比较常规，遍历输入字符串，数字字符，字母字符分别记录到两个不同数组，两个数组长度区别>1时，则明显不符题意，返回空字符串，否则则两个数组分别遍历组装字符串返回。
