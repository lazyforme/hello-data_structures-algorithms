package solution

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	sum := make([]int, length)
	for i := 0; i < length; i++ {
		if 0 == i {
			sum[i] += nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
	}
	result := NumArray{sum}
	return result
}

func (this *NumArray) SumRange(i int, j int) int {
	if 0 == i {
		return this.sum[j]
	} else {
		return this.sum[j] - this.sum[i-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

/**
    res := NumArray{
        acc: make([]int, len(nums)),
    }
    var acc int
    for i, v := range nums {
        acc = acc + v
        res.acc[i] = acc
    }
    return res
}
*/
