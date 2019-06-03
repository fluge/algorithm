package leetcode

/*
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
*/

func FindLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max := 0
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] { //表示不连续递增
			if max < i-index {
				max = i - index
			}
			index = i
		}
		if i == len(nums)-1 { //表示一直连续递增
			if max < i-index+1 {
				max = i - index + 1
			}
		}
	}
	return max
}
