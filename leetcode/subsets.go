package leetcode

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
//递归的实现
func SubSets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	res := SubSets(nums[1:])
	ll := len(res)
	for i := 0; i < ll; i++ {
		res = append(res, append([]int{nums[0]}, res[i]...))
	}
	return res
}
