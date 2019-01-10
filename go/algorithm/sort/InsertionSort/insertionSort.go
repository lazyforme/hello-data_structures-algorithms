package insertSort

func insertSort(array []int) []int {
	length := len(array)
	index := 1
	for i := 1; i < length; i++ {
		for j := 0; j < index; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
		index += 1
	}
	return array
}

func binarySearchInsertSort(array []int) []int {
	length := len(array)
	for i := 1; i < length; i++ {
		index := binarySearch(array, i-1, array[i])
		if i != index {
			temp := array[i]
			for j := i - 1; j >= index; j-- {
				array[j+1] = array[j]
			}
			array[index] = temp
		}
	}
	return array
}

func binarySearch(array []int, max int, key int) int {
	low, high := 0, max
	for low <= high {
		mid := low + (high-low)/2
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
