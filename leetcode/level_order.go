package leetcode

/*
二叉树的层次遍历

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

给定二叉树: [3,9,20,null,null,15,7],

  	3
   / \
  9  20
    /  \
   15   7



[
  [3],
  [9,20],
  [15,7]
]


思路： 借助队列实现,有一个问题是怎么记住层级。。尝试使用傀儡节点 ，，使用递归解法
*/

func LevelOrder(root *TreeNode) [][]int {
	list := make([]*TreeNode, 0)
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	//加入头节点
	list = append(list, root)
	//加入傀儡节点（使用-1 作为傀儡节点）
	list = append(list, &TreeNode{
		Val: MinInt,
	})

	tmp := make([]int, 0)
	//表示不为空
	for len(list) > 0 {
		if list[0].Val == MinInt && list[0].Right == nil && list[0].Left == nil {
			//遍历到傀儡节点了，就在队列里面插入一个傀儡节点
			list = append(list, &TreeNode{
				Val: MinInt,
			})
			if len(tmp) > 0 {
				res = append(res, tmp)
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
	return res
}
