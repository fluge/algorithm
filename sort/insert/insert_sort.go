package insert

/**
插入排序(结果升序)
1.最好情况：序列已经是升序排列，在这种情况下，需要进行的比较操作需（n-1）次。后移赋值操作为0次。即O(n)
2.最坏情况：序列已经是降序排列，那么此时需要进行的比较共有n(n-1)/2次。后移赋值操作是比较操作的次数加上 (n-1）次。即O(n^2)

总体时间复杂度 O(n^2) 原地排序 O(1)
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
