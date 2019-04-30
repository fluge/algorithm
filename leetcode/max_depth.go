package leetcode

/*
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

    3
   / \
  9  20
    /  \
   15   7

返回它的最大深度 3 。

*/
//递归实现
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	right := MaxDepth(root.Right) + 1
	left := MaxDepth(root.Left) + 1
	if right > left {
		return right
	}
	return left
}

//非递归实现 ,利用层次遍历
func MaxDepth2(root *TreeNode) int {
	//return len(LevelOrder(root))
	if root == nil {
		return 0
	}
	list := make([]*TreeNode, 0)
	//加入头节点
	list = append(list, root)
	//加入傀儡节点（使用-1 作为傀儡节点）
	list = append(list, &TreeNode{
		Val: MinInt,
	})

	tmp := make([]int, 0)
	max := 0
	//表示不为空
	for len(list) > 0 {
		if list[0].Val == MinInt && list[0].Right == nil && list[0].Left == nil {
			//遍历到傀儡节点了，就在队列里面插入一个傀儡节点
			list = append(list, &TreeNode{
				Val: MinInt,
			})
			if len(tmp) > 0 {
				max++
				//重新初始化一个tmp,表示另一层
				tmp = make([]int, 0)
			} else {
				break
			}
		} else {
			//不是傀儡节点，就把左右节点分别入队
			tmp = append(tmp, list[0].Val)
			//有左右节点，先从做节点入队，然后是有节点
			if list[0].Left != nil {
				list = append(list, list[0].Left)
			}
			if list[0].Right != nil {
				list = append(list, list[0].Right)
			}
		}
		//把当前元素出队
		list = list[1:]
	}
	return max
}
