package leetcode

/*
142. 环形链表 II

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/

//使用hash 表的方式是能解需要记录第一次成环的地方,应该还有其他解法
func detectCycle(head *ListNode) *ListNode {
	nodeMap := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

//使用快慢指针的方式做降低空间复杂度
//思路就是：先通过快慢指针找到是否是环。如果找到了环，然后在更节点遍历回来
func detectCycles(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { //表示有环，这个时候
			for pre != nil {
				if slow == pre {
					return pre
				}
				slow = slow.Next
				pre = pre.Next
			}
		}
	}
	return nil
}
