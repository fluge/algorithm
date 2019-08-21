package leetcode

/*
给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

输入:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出: 2
说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回2。

输入:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出: 1
说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1


思路：使用dfs的方式去做这件事情，把进过的能变为朋友圈的都置0 然后算个数

*/

func findCircleNum(M [][]int) int {
	visit := make([]int, len(M))
	ct := 0
	for i := range M {
		if visit[i] == 0 {
			dfsCircleNum(M, i, visit)
			ct++
		}
	}
	return ct
}
func dfsCircleNum(M [][]int, i int, visit []int) {
	for j, value := range M[i] {
		if visit[j] == 0 && value == 1 {
			visit[j] = 1
			dfsCircleNum(M, j, visit)
		}
	}
}
