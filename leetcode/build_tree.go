package leetcode

/*
根据一棵树的前序遍历与中序遍历构造二叉树。
你可以假设树中没有重复的元素。


前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]

	3
   / \
  9  20
    /  \
   15   7

前序：根  --> 左 ---> 右
中序：左  --> 根 ---> 右
*/

//递归实现
func buildTree(preorder []int, inorder []int) *TreeNode {
	//使用前序找根节点，然后通过中序找左右节点
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}
	k := 0
	for i, v := range inorder {
		if v == preorder[0] {
			k = i
			break
		}
	}
	node.Left = buildTree(preorder[1:k+1], inorder[:k])
	node.Right = buildTree(preorder[k+1:], inorder[k+1:])
	return node
}
