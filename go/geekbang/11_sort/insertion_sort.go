package sort

func insertionSort(input []int) []int {
	arr := input
	length := len(input)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
	return arr
}

func insertionSort_2(input []int) []int {
	arr := input
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func insert_sort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}

func bubble_sort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return input
}
