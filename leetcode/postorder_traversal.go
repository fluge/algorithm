package leetcode

/*
给定一个二叉树，返回它的 后序 遍历。

   1
    \
     2
    /
   3

输出: [3,2,1]

后序遍历：左 --> 右 --> 根节点
*/

//递归后序
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	if root.Left != nil {
		res = append(res, PostOrderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, PostOrderTraversal(root.Right)...)
	}
	res = append(res, root.Val)
	return res
}

/*
非递归：使用双栈法
*/
func PostOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//初始化一个栈
	stack := make([]*TreeNode, 0)
	//根节点入栈
	stack = append(stack, root)
	//结果集
	res := make([]int, 0)
	for len(stack) > 0 {
		//取栈顶元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(res) == 0 {
			res = append(res, node.Val)
		} else {
			//向slice 的头插入对应的值（类似双栈里面的第二个栈）
			res = append(res[:1], res[0:]...)
			res[0] = node.Val
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return res
}

////第三种思路
//func PostOrderTraversal3(root *TreeNode) []int {
//
//}
