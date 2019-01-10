# 795. Number of Subarrays with Bounded Maximum

## description

We are given an array `A` of positive integers, and two positive integers `L` and `R` (`L <= R`).

Return the number of (contiguous, non-empty) subarrays such that the value of the maximum array element in that subarray is at least `L` and at most `R`.

```
Example :
Input: 
A = [2, 1, 4, 3]
L = 2
R = 3
Output: 3
Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].

```

## solution

理解题意，属于数组A的连续，非空的子数组，其中该子数组的最大元素不能小于L，最大元素不能大于R。重点是子数组中的元素必须是数组A中连续的若干个元素，子数组中的最大元素必须落在区间[L,R]中。

首先，我们可以算出数组A所有的连续、非空子数组。

[2],[2,1],[2,1,4],[2,1,4,3]

[1],[1,4],[1,4,3]

[4],[4,3]

[3]

数组A具有4+3+2+1=10个连续、非空子数组。在这基础上再加上子数组中最大元素必须落在区间[L,R]上。[n1,n2,n3,....,nk]，假设nk(第一个)>R，R>=ni(第一个)>=L，则[ni,..nk-1]为满足条件的最长的连续子数组。

当只有ni>=L时，则子数组的个数为k-1-i+1

则由此即可得出此题的解法



