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

////非递归
//func postorderTraversal(root *TreeNode) []int {
//
//}
