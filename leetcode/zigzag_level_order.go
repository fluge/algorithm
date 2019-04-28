package leetcode

/**
二叉树的锯齿形层次遍历

给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

   	3
   / \
  9  20
    /  \
   15   7


[
  [3],
  [20,9],
  [15,7]
]

*/

func ZigzagLevelOrder(root *TreeNode) [][]int {
	//使用数组当队列，然后进行遍历
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

	index := true
	//表示不为空
	for len(list) > 0 {
		if list[0].Val == MinInt && list[0].Right == nil && list[0].Left == nil {
			//遍历到傀儡节点了，就在队列里面插入一个傀儡节点
			list = append(list, &TreeNode{
				Val: MinInt,
			})
			if len(tmp) > 0 {
				if index {
					//把整个数组反转
					index = false
				} else {
					index = true
				}
				res = append(res, tmp)
				//重新初始化一个tmp,表示另一层
				tmp = make([]int, 0)
			} else {
				break
			}
		} else {
			//不是傀儡节点，就把左右节点分别入队
			tmp = append(tmp, list[0].Val)

			if index {
				tmp := list[0].Left
				list[0].Left = list[0].Right
				list[0].Right = tmp
			}

			if list[0].Left != nil {
				list = append(list, list[0].Left)
			}
			//有左右节点，先从做节点入队，然后是有节点
			if list[0].Right != nil {
				list = append(list, list[0].Right)
			}

		}
		//把当前元素出队
		list = list[1:]
	}
	return res
}
