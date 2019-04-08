package leetcode

/*
给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。

返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
*/

/*
 看示例的意思是，使用翻转做排序，记录每次翻转的过程返回。
 使用暴力解法的话就是每次找到最大值，然后通过最多两次的翻转到最后面（类似冒泡排序），用递归解，边界条件是最后数组的长度大于1
*/

func PancakeSort(A []int) []int {
	res := make([]int, 0)
	for {
		arr, r := overTurn(A)
		res = append(res, r...)
		if len(arr) > 1 {
			A = arr[:len(arr)-1]
		} else {
			break
		}
	}

	return res
}

//找出最大值，并返回翻转后的数组
func overTurn(arr []int) ([]int, []int) {
	i := 0
	max := -1
	//找到最大值
	for k, v := range arr {
		if v > max {
			max = v
			i = k
		}
	}

	if i == len(arr)-1 {
		return arr, []int{}
	}

	//做两次翻转操作
	//第一次把最大值翻转到最前面
	arr = turn(arr, i)
	//第二次把最大值翻转到最最后面
	arr = turn(arr, len(arr)-1)
	//输出结果
	return arr, []int{i + 1, len(arr)}
}

func turn(arr []int, j int) []int {

	for i := 0; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
