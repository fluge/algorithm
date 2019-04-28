package leetcode

/**
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]

对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。


比较经典的深度优先

*/

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i, v := range grid {
		for in, value := range v {
			if value > 0 {
				temp := Dfs(grid, i, in)
				if temp > maxArea {
					maxArea = temp
				}
			}
		}
	}
	return maxArea
}
func Dfs(grid [][]int, i, j int) int {
	area := 0
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] > 0 {
		grid[i][j] = 0
		area = 1 + Dfs(grid, i-1, j) + Dfs(grid, i+1, j) + Dfs(grid, i, j+1) + Dfs(grid, i, j-1)
	}
	return area
}
