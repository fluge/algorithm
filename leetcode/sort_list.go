package leetcode

/**
148 . 排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
使用各种数组的经典排序，比如快排，归并
*/

//使用归并排序,这个和数组的归并不一样的是，链表不能随机遍历，需要找到中间节点 ，使用两个指针，一个步长为1 一个步长为2。当第二个遍历完，就找到中间节点
func mergeSortList(head *ListNode) *ListNode {
	//遍历到节点只剩下一个
	if head == nil || head.Next == nil {
		return head
	}
	p, q := head, head //分别为快慢指针
	for p.Next != nil && p.Next.Next != nil {
		p = p.Next.Next
		q = q.Next
	}
	l1 := head
	l2 := q.Next
	q.Next = nil

	left := mergeSortList(l1)
	right := mergeSortList(l2)
	return mergeTwoLists(left, right)
}

//// 使用快速排序 因为不能和数组一样从可以
//func quickSortList(head *ListNode) *ListNode {
//
//}
