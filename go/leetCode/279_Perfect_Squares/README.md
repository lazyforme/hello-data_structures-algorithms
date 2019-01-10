# 279. Perfect Squares

## description

Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

**Example 1:**
```
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
```

**Example 2:**
```
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
```

## solution

平方数组如[1,4,9,16,25,36,...]

dp[i]定义为整数i当前的平方数之和的最小个数。则dp[i]=min(dp[i-square]+1),square为平方数。
e.g: dp[12]=min(dp[12-1]+1, dp[12-4]+1, dp(12-9)+1)