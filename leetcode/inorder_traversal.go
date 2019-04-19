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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	lret := inorderTraversal(root.Left)
	if len(lret) > 0 {
		ret = append(ret, lret...)
	}
	ret = append(ret, root.Val)
	rret := inorderTraversal(root.Right)
	if len(rret) > 0 {
		ret = append(ret, rret...)
	}
	return ret
}
