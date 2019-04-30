package leetcode

/*
144. 二叉树的前序遍历
*/

//递归解法
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, PreOrderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, PreOrderTraversal(root.Right)...)
	}
	return res
}

//非递归解法
func PreOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	//把根节点入栈
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}
