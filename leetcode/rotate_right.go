package leetcode

import (
	"fmt"
)

/*
61.旋转链表
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

思路： 每次把最后一个节点移到最开始，然后根据k 进行遍历

*/

//把链表作为环，然后根据 k 找到头节点，然后进行切分
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := head
	length := 0
	//先成环，然后根据k 去遍历
	for root != nil {
		length++
		if root.Next == nil {
			root.Next = head
			head = root
			break
		}
		root = root.Next
	}

	// 找到 环的头节点
	n := k % length

	fmt.Println(n, head.Val)
	//return head

	for ; n >= 1; n-- {
		head = head.Next
		if n == 1 {
			node := head.Next
			head.Next = nil
			head = node
			break
		}
	}

	return head
}

//可以作为解，但是时间复杂度太高
func rotateRights(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head
	length := 0
	// 对 k 进行取 余操作，缩小后面遍历的范围
	for tmp != nil {
		length++
		if tmp.Next == nil {
			k = k % length
		}
		tmp = tmp.Next
	}

	root := head
	for ; k > 0; k-- {
		head = root
		//每次都进行遍历，把 最后一个节点
		for head != nil && head.Next != nil {
			if head.Next.Next == nil {
				node := head.Next
				head.Next = nil
				node.Next = root
				root = node
				break
			}
			head = head.Next
		}
	}
	return root
}
