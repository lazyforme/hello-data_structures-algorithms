# 213. House Robber II

## description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

```
Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
```

```
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
```

## solution

由于店铺排列是一个圆，要求不变依然是不能抢劫相邻的两个店铺。因为是一个圆，所以收尾相连，抢劫了shop[0]就不能抢劫shop[n-1],抢劫了shop[1]就不能抢劫shop[0],抢劫了shop[2]就不能抢劫shop[1]...按此规律，所以就可以把圆形店铺链分为[0, n-2]和[1, n-1]这两类店铺排列形式。只需要分别算出这个两类的值，再取其最大值即为所求。

dp[i] = max(dp[i-2],dp[i-3]) + nums[i]

原理依旧是斐波那契数列