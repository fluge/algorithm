package leetcode

/**
5. 最长回文子串

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。


输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

输入: "cbbd"
输出: "bb"

思路：寻找回文的中间值，然后两变遍历，之后变指针，直到结尾(需要考虑两种情况) (中心扩散法)

*/

// 中心扩散
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	max := ""
	for i := 0; i < len(s)-1; i++ {
		m := sub(i, i, s)
		n := sub(i, i+1, s)
		if len(m) > len(max) {
			max = m
		}
		if len(n) > len(max) {
			max = n
		}
	}
	return max
}

func sub(a, b int, str string) string {
	max := ""
	//找到该指针的回文值 (单指的回文)
	for a >= 0 && b < len(str) {
		if str[a] == str[b] {
			a--
			b++
		} else {
			break
		}
	}
	if len(str[a+1:b]) > len(max) {
		max = str[a+1 : b]
	}
	return max
}
