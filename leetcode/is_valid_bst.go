package leetcode

/*
98. 验证二叉搜索树

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

输入:
    2
   / \
  1   3
输出: true

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

思路1：最简单的还是递归，使用极大值和极小值，做递归判断
思路2：使用中序遍历（左，根，右），如果是bst，这个数据就直接是有序的
*/

func IsValidBST(root *TreeNode) bool {
	return isValidBSTImpl(root, MinInt, MaxInt)
}

func isValidBSTImpl(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && isValidBSTImpl(root.Left, min, root.Val) && isValidBSTImpl(root.Right, root.Val, max)
}

//使用中序遍历的方式去做，只用一个栈就可以完成
func IsValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	isBst := true
	pre := MaxInt
	//把根节点入栈
	for root != nil || len(stack) > 0 {
		//把左子树一直入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//对左子树出栈
		root = stack[len(stack)-1]
		if pre != MaxInt && pre >= root.Val {
			isBst = false
			break
		}
		pre = root.Val
		stack = stack[:len(stack)-1]
		//有节点
		root = root.Right
	}
	return isBst
}
