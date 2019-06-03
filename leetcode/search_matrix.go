package leetcode

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

给定 target = 5，返回 true。

给定 target = 20，返回 false。

思路：从右上角开始找，该值比目标值大就下移，比目标值小就左移
*/

func SearchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	if height == 0 {
		return false
	}
	width := len(matrix[0])
	if height == 0 {
		return false
	}

	x := 0
	y := width - 1
	for x < height && y >= 0 {
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target { // 左移
			y--
		} else {
			x++
		}
	}
	return false

}
