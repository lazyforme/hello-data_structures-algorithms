# 771. Jewels and Stones

## description

You're given strings J representing the types of stones that are jewels, and S representing the stones you have. Each character in S is a type of stone you have. You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A"

**Note:**

- `S`and`J`will consist of letters and have length at most 50.

- The characters in`J`are distinct.

## solution

理解题意，题目给定两个字符串J、S，要统计出在字符串S中出现过字符串J中每个字符的个数。例如，J="aA"，S="Aaasd"，J中有'a','A'两个字符，这两个字符在S中出现了3次。

解法一：

首先想到的就是利用循环，或者循环字符串J，或者循环字符串S，只需要算出在S中出现过字符j的个数即可。每个编程语言的是实现方式各不相同，但主题思想都是一样的。
