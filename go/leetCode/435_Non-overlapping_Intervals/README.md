# 435. Non-overlapping Intervals

## description

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

### Note:

You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

### Example 1:

```
Input: [ [1,2], [2,3], [3,4], [1,3] ]

Output: 1

Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
```

### Example 2:

```
Input: [ [1,2], [1,2], [1,2] ]

Output: 2

Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
```

### Example 3:

```
Input: [ [1,2], [2,3] ]

Output: 0

Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

## solution

计算最多能组成的不重叠区间个数，然后用区间总个数减去不重叠区间的个数。

在每次选择中，区间的结尾最为重要，选择的区间结尾越小，留给后面的区间的空间越大，那么后面能够选择的区间个数也就越大。

按区间的结尾进行排序，每次选择结尾最小，并且和前一个区间不重叠的区间。

使用贪心算法，类似于`455`，要事先将数组排序，对已排序好的数组进行局部最优比对，从而得到全局最优。但并不是所有情况都适用于贪心算法，即所有的局部最优加起来并不等于全局最优。