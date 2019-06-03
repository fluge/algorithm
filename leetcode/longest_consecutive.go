package leetcode

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。


输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

思路：时间复杂度为O(n) 的话 需要一次遍历就完成，不能使用排序来做，排序最少需要的 n*log n来做
*/

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	tmp := make(map[int]bool, 0)
	//先把所有的放到map 里面
	for _, v := range nums {
		tmp[v] = true
	}
	//然后开始遍历map
	for k := range tmp {
		//判断 k-1 是否在map 里面,k 就是最小的连续
		if ok, _ := tmp[k-1]; !ok {
			i := k + 1
			//判断k+1 是否存在
			in, _ := tmp[i]
			for in {
				i++
				if max < i-k {
					max = i - k
				}
				in, _ = tmp[i]
			}
		}
	}
	return max
}
