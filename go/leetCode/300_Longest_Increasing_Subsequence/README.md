# 300.Longest Increasing Subsequence

## description

Given an unsorted array of integers, find the length of longest increasing subsequence.

**Example**:

```
Input: [10,9,5,2,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

**Note:**

* There may be more than one LIS combination, it is only necessary for you to return the length.
* Your algorithm should run in O(n2) complexity.

Follow up: Could you improve it to O(n log n) time complexity?

## solution

动态规划，定义一个数组 dp 存储最长递增子序列的长度，dp[n] 表示以 Sn 结尾的序列的最长递增子序列长度。

dp[n] = max{ 1, dp[i]+1 | Si < Sn && i < n} 。

对于一个长度为 N 的序列，最长递增子序列并不一定会以 SN 为结尾，因此 dp[N] 不是序列的最长递增子序列的长度，需要遍历 dp 数组找出最大值才是所要的结果，max{ dp[i] | 1 <= i <= N} 即为所求。

