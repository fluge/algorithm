package chapter_4

import (
	"github.com/fluge/algorithm/leetcode"
)

/*
最大子数组
*/

//使用分治法 时间复杂读为 O(n*log n)
func FindMaximumSubArray(A []int, low, high int) (arrLow, arrHigh, arrSum int) {
	if low == high {
		return low, high, A[low]
	}

	mid := (low + high) / 2
	//只有3种情况
	//情况1：最大子序列出现在A[low]~A[mid]
	//情况2：最大子序列出现在A[mid+1]~A[high]
	//情况3：最大子序列包含A[mid],A[mid+1]
	leftLow, leftHigh, leftSum := FindMaximumSubArray(A, low, mid)
	rightLow, rightHigh, rightSum := FindMaximumSubArray(A, mid+1, high)
	corssLow, corssHigh, corssSum := findMaximumCorssSubArray(A, low, mid, high)

	if leftSum >= rightSum && leftSum >= corssSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= corssSum {
		return rightLow, rightHigh, rightSum
	}
	return corssLow, corssHigh, corssSum
}

func findMaximumCorssSubArray(A []int, low, mid, high int) (arrLow, arrHigh, arrSum int) {
	leftSum := leetcode.MinInt
	sum := 0
	for i := mid; i >= low; i-- {
		sum = sum + A[i]
		if sum > leftSum {
			leftSum = sum
			arrLow = i
		}
	}

	rightSum := leetcode.MinInt
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum = sum + A[i]
		if sum > rightSum {
			rightSum = sum
			arrHigh = i
		}
	}
	return arrLow, arrHigh, leftSum + rightSum
}
