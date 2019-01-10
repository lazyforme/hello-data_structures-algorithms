# 167. Two Sum II - Input array is sorted

## description

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

### Note:

* Your returned answers (both index1 and index2) are not zero-based.

* You may assume that each input would have exactly one solution and you may not use the same element twice.

### Example:

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

## solution

题中给出的是升序数组，那么第一个元素是最小的，最后一个元素是最大的。我们可以使用两个标记指针low和high，分别指向第一个和最后一个，两个指针慢慢相互靠拢。当两个指针所属元素之和大于target时，说明high指向的元素大了，high--；反之，low++。