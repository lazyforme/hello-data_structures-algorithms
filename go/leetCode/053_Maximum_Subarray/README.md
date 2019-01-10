# 53. Maximum Subarray

## description

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Follow up:**

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## solution

由于题目要求连续数组（最少包含一个元素）内元素之和最大，关键是连续，需要一个标记位来记录当前的最大值，然后循环该数组。