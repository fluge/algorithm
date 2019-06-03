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
func BuildTree(preorder []int, inorder []int) *TreeNode {
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
	node.Left = BuildTree(preorder[1:k+1], inorder[:k])
	node.Right = BuildTree(preorder[k+1:], inorder[k+1:])
	return node
}

//根据中序和后序来构造二叉树，原理和前序和中序是一样的。通过后序找根节点
func BuildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val: postorder[len(postorder)-1], //通过后序的最后一个节点找到对应的根节点
	}
	if len(postorder) == 1 || len(inorder) == 1 {
		return node
	}

	k := 0
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			k = i
			break
		}
	}
	node.Left = BuildTree2(inorder[:k], postorder[:k])
	node.Right = BuildTree2(inorder[k+1:], postorder[k:len(postorder)-1])
	return node
}
