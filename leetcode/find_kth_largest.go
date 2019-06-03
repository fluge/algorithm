package leetcode

/*
  数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/
//思路是全局排序，然后遍历选第k个元素，可以用最小堆做，把前k个元素建堆，然后最后遍历输出就行
func findKthLargest(nums []int, k int) int {
	//使用快排进行排序
	sortQuick(nums)
	for i, v := range nums {
		if i+1 == k {
			return v
		}
	}
	return -1
}

func sortQuick(nums []int) {
	if len(nums) <= 1 {
		return
	}
	tmp, i := nums[0], 1
	head, tail := 0, len(nums)-1
	for tail > head {
		//左边第一个元素大于选定值，直接放到最后
		if nums[i] > tmp {
			nums[i], nums[head] = nums[head], nums[i]
			//只要保证右边的逗比选定值大就行
			head++
			i++
		} else {
			//小于选定值，移到把分水岭的左边
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}
	sortQuick(nums[:head])
	sortQuick(nums[head+1:])
}
