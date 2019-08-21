package leetcode

/*
141. 环形链表
给定一个链表，判断链表中是否有环。

两种解法：快慢指针和哈希表
*/

//使用hash 解法 空间和时间都是 O(n)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

//采用快慢指针进行操作
func hasCycles(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
