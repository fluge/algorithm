package leetcode

/*
第k个排列
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。

输入: n = 3, k = 3
输出: "213"

输入: n = 4, k = 9
输出: "2314"

思路：就是把n个排列给输出出来。然后输出第k个元素
*/

//func getPermutation(n int, k int) string {
//
//}