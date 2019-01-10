package sort

func bubbleSort(input []int) []int {
	arr := input
	lenght := len(arr)
	if lenght <= 1 {
		return arr
	}
	for i := 0; i < lenght; i++ {
		flag := false
		for j := 0; j < lenght-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
