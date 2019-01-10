package bubbleSort

func bubbleSort(array []int) []int {
	length := len(array)

	for i := 0; i < length-1; i++ {
		isChange := false
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}

	return array
}
