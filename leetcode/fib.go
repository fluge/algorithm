package leetcode

/*
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
*/

//暴力解法，递归
func Fib(N int) int {
	if N > 2 {
		return Fib(N-1) + Fib(N-2)
	}
	if N == 0 {
		return 0
	}
	return 1
}

//todo 优化解法
