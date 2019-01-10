# 540. Single Element in a Sorted Array

## description

Given a sorted array consisting of only integers where every element appears twice except for one element which appears once. Find this single element that appears only once.

```
Example 1:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
```

```
Example 2:
Input: [3,3,7,7,10,11,11]
Output: 10
```

**Note**: Your solution should run in O(log n) time and O(1) space.

## solution

二分法。

通过题目可以得出该数组的性质：

* 该数组的长度为奇数；

* 在single element出现之前的元素，其都是两两相等的，nums[0]=nums[1], nums[2]=nums[3]...，即nums[偶数]==nums[奇数]

* 在single element出现之后的元素，其都是两两相等的，由于中间插入了一个single element，打乱了原有的出顺序，变成了nums[5]=nums[6], nums[7]=nums[8]...，即nums[奇数]==nums[偶数]

我们可以利用这一性质，结合二分法来解决。