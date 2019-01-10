package sort

func QuickSort(input []int) {
	if len(input) <= 1 {
		return
	}
	pivot := Partition(input)
	QuickSort(input[:pivot])
	QuickSort(input[pivot+1:])
}

func Partition(input []int) (index int) {
	pivot := input[len(input)-1]
	i := 0
	for j := 0; j < len(input)-1; j++ {
		if input[j] < pivot {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}
	input[i], input[len(input)-1] = input[len(input)-1], input[i]
	return i
}
