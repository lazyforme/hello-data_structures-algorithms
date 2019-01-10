# 001. Two Sum

## description

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

## solution

解法一：

最先想到的思路是对数组使用两次循环，nums[i] + nums[i + n] (注:n=1...nums.length) == target，其中i和i+n，即为所求的数组索引，但是该解法的时间复杂度是O(n*n)，并不推荐。其中有一个讨巧的方法，我们不必计算nums[i]+nums[j]=target，在第一次循环数组nums时，求得target-nums[i]=tmp，只要在第二次循环数组nums时，找到元素值等于tmp即可。

解法二：

使用map的特点，key-value形式，我们可以将数组nums的值作为map的key，将index作为map的value。找到map[target-nums[i]]的value，该value和i即为所求的index。
