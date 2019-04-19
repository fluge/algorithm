package leetcode

import "fmt"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
//暴力解法 用hashmap存
func LengthOfLongestSubstring(s string) int {
	strMap := make(map[int32]int, 0)
	start := 0
	max := -1
	for k, v := range s {
		//如果值存在表示存在
		if value, ok := strMap[v]; ok {
			//开始指针
			if start <= value {
				start = value + 1
			}

		}
		strMap[v] = k
		if k-start > max {
			max = k - start
		}
		fmt.Println(start, k)
	}
	//从0开始的index 要加一
	return max + 1
}
