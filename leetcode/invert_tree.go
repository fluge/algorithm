package leetcode

/*
226. 翻转二叉树

翻转一棵二叉树。

 	 4
   /   \
  2     7
 / \   / \
1   3 6   9


     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/

//递归版
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil && (root.Left != nil || root.Right != nil) {
		root.Right, root.Left = InvertTree(root.Left), InvertTree(root.Right)
	}
	return root
}

//非递归版， 使用栈进行实现
func InvertTree2(root *TreeNode) *TreeNode {
	//直接返回根节点
	if root == nil || (root.Right == nil && root.Left == nil) {
		return root
	}
	//初始化一个栈信息
	stack := make([]*TreeNode, 0)
	//把根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		//节点出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//把node的左右节点互换
		node.Left, node.Right = node.Right, node.Left
		//把左右节点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
