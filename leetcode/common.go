package leetcode

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

//主要还是golang中最小int 的定义
const MinInt = -MaxInt - 1

//树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//链表
type ListNode struct {
	Val  int
	Next *ListNode
}
