package leetcode

/*
搜索旋转排序数组

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。


输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4

思路：O(log n)的查找，基本能想到的是二分查找，然后这个就是先找到旋转点，然后以旋转点进行二分查找
*/

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums) - 1
	divisionIndex := findDivision(nums)
	if divisionIndex == 0 || divisionIndex == -1 {
		//非旋转排序数组
		return findTarget(nums, 0, n, target)
	}

	res := findTarget(nums, 0, divisionIndex-1, target)
	if res != -1 {
		return res
	}
	return findTarget(nums, divisionIndex, n, target)
}

//找到分割点
func findDivision(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1

		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func findTarget(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
