package leetcode

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。


输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

//逆序链表 能想到的暴力就是把链表转换为int 然后相加后在存如一个新链表
 忽略了大数相加容易溢出

在短链表进行补0，然后遍历相加
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := carry + x + y

		carry = sum / 10
		head.Next = &ListNode{
			Val: sum % 10,
		}
		head = head.Next
	}
	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}
	return root.Next
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := &ListNode{}
//	p, q := l1, l2
//	curr := head
//	carry := 0
//	for p != nil || q != nil {
//		x := 0
//		if p != nil {
//			x = p.Val
//		}
//		y := 0
//		if q != nil {
//			y = q.Val
//		}
//		sum := carry + x + y
//		carry = sum / 10
//		curr.Next = &ListNode{Val: sum % 10}
//		curr = curr.Next
//		if p != nil {
//			p = p.Next
//		}
//		if q != nil {
//			q = q.Next
//		}
//	}
//	if carry > 0 {
//		curr.Next = &ListNode{Val: carry}
//	}
//	return head.Next
//}
