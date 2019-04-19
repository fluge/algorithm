package leetcode

import (
	"fmt"
	"testing"
)

func TestPancakeSort(t *testing.T) {
	fmt.Println(PancakeSort([]int{3, 2, 4, 1}))
}

func TestFib(t *testing.T) {
	fmt.Println(Fib(4))
}

func TestThirdMax(t *testing.T) {
	fmt.Println(ThirdMax([]int{3, 2, 1}))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcbb"
	fmt.Println(LengthOfLongestSubstring(str))
}

func TestLongestCommonPrefix(t *testing.T) {
	str := []string{"aa", "a", "aaa"}
	fmt.Println(LongestCommonPrefix(str))
}
