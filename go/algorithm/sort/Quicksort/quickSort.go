package quickSort

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	lessList, greaterList := []int{}, []int{}
	pivot := array[0] //选取基准值
	for i := 1; i < len(array); i++ {
		if array[i] <= pivot {
			lessList = append(lessList, array[i])
		} else {
			greaterList = append(greaterList, array[i])
		}
	}
	return append(append(quickSort(lessList), pivot), quickSort(greaterList)...)
}

func partition(array []int, left int, right int, pivotIndex int) int {
	pivotValue := array[pivotIndex]
	array[right], array[pivotIndex] = array[pivotIndex], array[right]
	storeIndex := left
	for i := left; i < right; i++ {
		if array[i] < pivotValue {
			array[i], array[storeIndex] = array[storeIndex], array[i]
			storeIndex += 1
		}
	}
	array[right], array[storeIndex] = array[storeIndex], array[right]
	return storeIndex
}

func quickSort_II(array []int, left int, right int) {
	if left < right {
		pivotIndex := left //选择基准值的index，这步可以优化
		pivotNewIndex := partition(array, left, right, pivotIndex)
		quickSort_II(array, left, pivotNewIndex-1)
		quickSort_II(array, pivotNewIndex+1, right)
	}
}

func quickSort_III(array []int) {
	if len(array) < 2 {
		return
	}
	low, high := 0, len(array)-1
	mid := array[0]
	for i := 1; i <= high; {
		if array[i] > mid {
			array[i], array[high] = array[high], array[i]
			high -= 1
		} else {
			array[i], array[low] = array[low], array[i]
			low += 1
			i++
		}
	}
	quickSort_III(array[:low])
	quickSort_III(array[low+1:])
}
