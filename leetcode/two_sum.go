package leetcode

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
*/
//暴力解法 时间复杂度是 O(n^2) 空间复杂度是O(1)
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1}
}

//空间换时间  时间和空间复杂度 都是O(n)
func TwoSum2(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for k, v := range nums {
		numMap[v] = k
	}

	for k, v := range nums {
		complement := target - v
		if va, ok := numMap[complement]; ok {
			if va == k {
				continue
			}
			return []int{k, va}
		}
	}
	return []int{1}
}
