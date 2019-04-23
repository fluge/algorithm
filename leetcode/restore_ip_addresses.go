package leetcode

import (
	"fmt"
	"strconv"
)

/*
复原IP地址

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

思路：
*/

func RestoreIpAddresses(s string) []string {
	return restore(s, 4)
}

func restore(s string, require int) []string {
	if require == 1 {
		if len(s) > 1 && '0' == s[0] {
			return []string{}
		}
		if v, _ := strconv.Atoi(s); v < 256 {
			return []string{s}
		}
		return []string{}
	}

	var r []string
	for i := 1; i < 4 && i <= len(s); i++ {
		prefix := s[:i]
		//fmt.Println("prefix:", prefix, " s:", s, " require:", require)
		if v, _ := strconv.Atoi(prefix); v < 256 {
			for _, j := range restore(s[i:], require-1) {
				fmt.Println("j", j, " prefix:", prefix)
				r = append(r, prefix+"."+j)
			}
		}
		if '0' == s[0] {
			break
		}
	}
	fmt.Println("r:", r)
	return r
}

func restoreIpAddresses(s string) []string {
	var res []string
	findIp(s, 0, 0, "", &res)
	return res
}

func findIp(s string, f int, idx int, ip string, res *[]string) {
	if idx == 3 {
		if len(s)-f <= 3 {
			if s[f] == '0' && (f != len(s)-1) {
				return
			}
			num, _ := strconv.Atoi(s[f:])
			if num < 256 {
				*res = append(*res, ip+strconv.Itoa(num))
			}
		}
	} else {
		for i := 1; i <= 3; i++ {
			if f+i >= len(s) {
				break
			}
			num, _ := strconv.Atoi(s[f : f+i])
			if num < 256 {
				findIp(s, f+i, idx+1, ip+strconv.Itoa(num)+".", res)
			}
			if s[f] == '0' && i == 1 {
				break
			}

		}
	}
}
