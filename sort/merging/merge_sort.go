package merging

/*
归并排序，主要采用分治法去处理

稳定排序  O(nlog n) 不是原地排序 O(n)
*/
//递归的方式
func SortMerge(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	//递归
	middle := len(data) / 2
	//不断地进行左右对半划分,类似与栈的动作，先进后出
	left := SortMerge(data[:middle])
	right := SortMerge(data[middle:])
	//对数组进行合并
	return merge(left, right)
}

//对两个数据做合并动作
func merge(left, right []int) []int {
	//用空间换时间,建立新的数组，不是就地排序
	result := make([]int, 0)
	i, j := 0, 0
	// 左右两个数组进行合并
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	//对两边数组的结尾进行处理，left和right 中有一个是0，直接进行append
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

//todo 非递归的方式
