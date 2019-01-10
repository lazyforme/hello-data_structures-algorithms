# [QuickSort](https://www.wikiwand.com/zh/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F)

## description

快速排序（英语：Quicksort），又稱劃分交換排序（partition-exchange sort），簡稱快排，一種排序算法，最早由東尼·霍爾提出。在平均狀況下，排序n個項目要{\displaystyle \ O(n\log n)}（大O符号）次比較。在最壞狀況下則需要{\displaystyle O(n^{2})}次比較，但這種狀況並不常見。事實上，快速排序{\displaystyle \Theta (n\log n)}通常明顯比其他演算法更快，因為它的內部循环（inner loop）可以在大部分的架構上很有效率地達成。

演算法
快速排序使用分治法（Divide and conquer）策略來把一個序列（list）分為兩個子序列（sub-lists）。

步驟為：

	1. 從數列中挑出一個元素，稱為「基準」（pivot），
	2. 重新排序數列，所有比基準值小的元素擺放在基準前面，所有比基準值大的元素擺在基準後面（相同的數可以到任何一邊）。在這個分割結束之後，該基準就處於數列的中間位置。這個稱為分割（partition）操作。
	3. 递归地（recursively）把小於基准值元素的子數列和大於基准值元素的子數列排序。

递归到最底部时，數列的大小是零或一，也就是已經排序好了。這個演算法一定會結束，因為在每次的迭代（iteration）中，它至少會把一個元素擺到它最後的位置去。