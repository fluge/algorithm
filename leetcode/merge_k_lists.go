package leetcode

/*
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
*/

//思路一：是和类似合并两个链表，进行合并，
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 1; i < len(lists); i++ {
		root := &ListNode{}
		h := root
		for head != nil || lists[i] != nil {
			if head == nil {
				h.Next = lists[i]
				break
			}
			if lists[i] == nil {
				h.Next = head
				break
			}

			if lists[i].Val > head.Val {
				h.Next = head
				head = head.Next
			} else {
				h.Next = lists[i]
				lists[i] = lists[i].Next
			}
			h = h.Next
		}

		head = root.Next
	}
	return head
}

//思路二 ： 思路一是一个个合并的，可以使用归并，在数组里面进行两两进行合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	//
	if len(lists) == 1 { //表示划分的数组的最小形式
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

//思路三： 就是利用优先级队列做事情
