package insert

/**
插入排序
*/

//升序和降序
func SortInsert(arr []int, desc bool) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && ((arr[j] > arr[j-1] && desc) || (arr[j] < arr[j-1] && !desc)); j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
