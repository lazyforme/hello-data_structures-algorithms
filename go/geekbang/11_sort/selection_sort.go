package sort

func selectionSort(input []int) []int {
	arr := input
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length-1; j++ {
			if arr[min] > arr[j+1] {
				min = j + 1
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
