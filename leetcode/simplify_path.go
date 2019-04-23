package leetcode

/*
简化路径

以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。

输入："/a/./b/../../c/"
输出："/c"

输入："/a//b////c/d//././/.."
输出："/a/b/c"

思路：/.. 指针向上移动一个， /.指针不动， 利用栈来实现
*/

func SimplifyPath(path string) string {
	arr := make([]string, 0)
	start := 0
	for k, v := range path {
		if v == 47 && start != k {
			arr = append(arr, string(path[start+1:k]))
			start = k
		}
		if k == len(path)-1 && start != k {
			arr = append(arr, string(path[start+1:k+1]))
		}
	}

	//初始化一个栈，用数组代替
	stack := make([]string, 0)
	//遍历来判断最后的路径
	for _, v := range arr {
		if v == "." || v == "" || v == "	" {
			continue
		}
		//类似pop
		if v == ".." {
			if len(stack) >= 1 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		//类似push
		stack = append(stack, v)
	}
	str := ""
	for _, v := range stack {
		str = str + "/" + v
	}
	if str == "" {
		str = "/"
	}
	return str
}
