# 303. Range Sum Query - Immutable

## description

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
```
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
```

## solution

定义sum[i]为[0, i]之和

[i, j]之和为sum[j] - sum[i-1]; e.g: [0, j] --> sum[j]; [1, j] --> sum[j] - sum[0]