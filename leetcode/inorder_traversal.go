package leetcode

/*
给定一个二叉树，返回它的中序 遍历。

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/

//中序是先 左，中，右
func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	lret := InOrderTraversal(root.Left)
	if len(lret) > 0 {
		ret = append(ret, lret...)
	}
	ret = append(ret, root.Val)
	rret := InOrderTraversal(root.Right)
	if len(rret) > 0 {
		ret = append(ret, rret...)
	}
	return ret
}

//中序遍历的非递归 --- 使用栈
func InOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	//把根节点入栈
	//stack = append(stack, root)
	for root != nil || len(stack) > 0 {
		//把左子树一直入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//对左子树出栈
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		//有节点
		root = root.Right
	}
	return res
}
