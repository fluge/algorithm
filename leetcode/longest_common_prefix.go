package leetcode

/*
最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

输入: ["flower","flow","flight"]
输出: "fl"
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	index := len(strs[0])
	str := ""
	for i := 1; i <= index; i++ {
		t := ""
		for k, v := range strs {
			if k == 0 {
				t = string(v[:i])
			} else {
				if len(v) < i || string(v[:i]) != t {
					return str
				}
			}
		}
		str = t
	}
	return str
}
