# 709. To Lower Case

## description

Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

## solution

题中要求是字符串的大小写转换，将字符串中的大写字母转换为小写字母。

解法一：

首先想到的是字母的ASCII编码，大写字母A~Z的编码为65~90，小写字母的编码为97~122。其中大小写字母的编码之差为32，也就是只要将字母转换成对应的编码，判断其的属于哪个范围，相应的加减32，即可得到其对应的大小写形式。

![image-20180722153111359](/Users/wcy/workspace/wcy/private/awesome-salary/wcy/leetCode/709_To_Lower_Case/ascii.png)

解法二：

构建一个以Key-Value形式的大小写字典库，循环传入的字符串，判断其大小写，再进行转换。