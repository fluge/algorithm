package leetcode

/**
  翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

*/
func ReverseWords(s string) string {
	//强制加一个_字符串 减少判断
	s = " " + s
	str := ""
	end := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		//如果遇到空格
		if string(s[i]) == " " {
			//做str 处理
			if i+1 != end {
				if str == "" {
					str = s[i+1 : end]
				} else {
					str = str + " " + s[i+1:end]
				}
			}
			//移动两个指针在一起
			end = i
		}
	}
	return str
}
