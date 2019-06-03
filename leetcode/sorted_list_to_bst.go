package leetcode

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

思路直接中dfs解
*/

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return dfs(res, 0, len(res)-1)
}

//
////找到中间节点，然后两边分别为左右节点
//func listDfs(head *ListNode) *TreeNode {
//	res
//	//找点链表的长度
//	for head.Next != nil {
//		length++
//		head = head.Next
//	}
//	l := length / 2
//	var listNode *ListNode
//	for head.Next != nil {
//		if l == 0 {
//			listNode =
//		}
//	}
//
//	node := &TreeNode{
//		Val: head.Val,
//		Left:listDfs(head),
//		Right:listDfs()
//	}
//	return node
//}
