package sort

func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) (res []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}
