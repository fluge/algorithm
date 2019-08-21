package leetcode

/**
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.

正常思路是 遍历两次，第一次找到整个列表的长度，第二次找到第n个节点 (把倒数变为正数)
如果是遍历一次，就是两个指针，间隔n 开始遍历，后一个节点遍历玩了，就删除第一个节点，然后进行输出就行
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root1, root2 := head, head
	length := 0
	for root1 != nil {
		length++
		root1 = root1.Next
	}
	carry := length - n

	if carry == 0 {
		return head.Next
	}
	length = 1
	for root2 != nil {
		//删除这个节点
		if length == carry {
			root2.Next = root2.Next.Next
		}
		length++
		root2 = root2.Next
	}

	return head
}

//遍历一次的实现
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	root1, root2 := head, head
	length := 1
	for root1 != nil {
		if length > n+1 {
			root2 = root2.Next
		}
		length++
		root1 = root1.Next
	}
	if length == n+1 {
		return head.Next
	}
	root2.Next = root2.Next.Next
	return head
}
