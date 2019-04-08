package leetcode

/*
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
*/

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

//主要还是golang中最小int 的定义
const MinInt = -MaxInt - 1

func ThirdMax(nums []int) int {

	a := MinInt
	b := MinInt
	c := MinInt
	for _, v := range nums {
		if v > a {
			c = b
			b = a
			a = v
		}
		if v < a && v > b {
			c = b
			b = v
		}

		if v > c && v < b {
			c = v
		}
	}
	if c != MinInt {
		return c
	}
	return a

}
