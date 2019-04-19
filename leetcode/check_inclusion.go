package leetcode

/**
字符串的排列

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串

思路：s1 的所以字母可以进行重新排列，s1建一个hash表。如果s2的固定长度全部都在hash表中表示包含
*/
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	strMap := make(map[int32]int, 0)
	for _, v := range s1 {
		if value, ok := strMap[v]; ok {
			//hash 表计数
			strMap[v] = value + 1
		} else {
			strMap[v] = 1
		}
	}
	//遍历s2
	for i := len(s1); i <= len(s2); i++ {
		//建立一个hash 表
		sMap := make(map[int32]int, 0)
		for _, v := range s2[i-len(s1) : i] {
			if value, ok := sMap[v]; ok {
				//hash 表计数
				sMap[v] = value + 1
			} else {
				sMap[v] = 1
			}
		}
		tmp := true
		//对比两个map
		for k, v := range sMap {
			if value, ok := strMap[k]; ok {
				//hash 表计数
				if value != v {
					tmp = false
					break
				}
			} else {
				tmp = false
				break
			}
		}
		if tmp {
			return true
		}
	}
	return false
}
