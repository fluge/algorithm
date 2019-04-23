package leetcode

import "sort"

/*
三数之和

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

注意：答案中不可以包含重复的三元组。


满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
//有一个[0 1 -1] 重复，需要对数组进行去重

思路：暴力方式：三个for循环，输出三个数之和，时间复杂度是O(n^3)
*/

func ThreeSum(nums []int) [][]int {
	arr := make([][]int, 0)
	t := map[[2]int]int{}

	//对nums 进行排序
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		for j := len(nums) - 1; j > i && nums[j] >= 0; j-- {
			target := nums[i] + nums[j]
			for k := i + 1; k < j; k++ {
				// 找到目标值
				if nums[k]+target == 0 {
					if _, ok := t[[2]int{nums[i], nums[j]}]; !ok {
						a, b, c := sortThree(nums[i], nums[k], nums[j])
						arr = append(arr, []int{a, b, c})
						t[[2]int{nums[i], nums[j]}] = nums[k]
						t[[2]int{nums[i], nums[k]}] = nums[j]
						t[[2]int{nums[k], nums[j]}] = nums[i]

					}
				}
			}
		}
	}
	return arr
}

func sortThree(x, y, z int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	if y > z {
		y, z = z, y
	}
	if x > y {
		x, y = y, x
	}
	return x, y, z
}
