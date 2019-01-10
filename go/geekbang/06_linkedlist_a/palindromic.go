package linkedlist_a

/** 如何判断字符串是回文 */
func palindromicString(input string) bool {
	length := len(input)
	if length < 2 {
		return true
	}
	head := 0
	tail := length - 1 - head
	for head < tail {
		if input[head] != input[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

/** 如何判断数组存储的字符串是回文 */
func palindromicArray(input []string) bool {
	length := len(input)
	if length < 2 {
		return true
	}
	head := 0
	tail := length - 1 - head
	for head < tail {
		if input[head] != input[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

/** 如何判断链表存储的字符串是回文 */
func palindromicLinkedList(input LinkedList) bool {
	length := input.size
	if length < 2 {
		return true
	}
	start := 0
	end := length - 1 - start
	for start < end {
		if input.Get(start) != input.Get(end) {
			return false
		}
		start++
		end--
	}
	return true
}
