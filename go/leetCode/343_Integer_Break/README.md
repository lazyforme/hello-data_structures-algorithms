# 343. Integer Break

## description

Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

**Example 1:**
```
Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
```

**Example 2:**

```
Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
```

**Note:** You may assume that n is not less than 2 and not larger than 58.

# solution

dp[i]为整数i当前最大分解数的乘积，可以动态规划分解为求dp[1],dp[2],dp[3],....,dp[i-1]，则dp[i]=max(dp[i],max(j×dp[i-j]),j×(i-j)).或者dp[i]=max(dp[j]×dp[i-j])