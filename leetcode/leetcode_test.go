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

func TestCheckInclusion(t *testing.T) {
	fmt.Println(CheckInclusion("ab", "ab"))
}

func TestReverseWords(t *testing.T) {
	fmt.Println(ReverseWords("hello   the   world   "))
}

func TestSimplifyPath(t *testing.T) {
	fmt.Println(SimplifyPath("/home/"))
	fmt.Println(SimplifyPath("/../"))
	fmt.Println(SimplifyPath("/home//foo/"))
	fmt.Println(SimplifyPath("/a/./b/../../c/"))
	fmt.Println(SimplifyPath("/a/../../b/../c//.//"))
	fmt.Println(SimplifyPath("/a//b////c/d//././/.."))
}

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(RestoreIpAddresses("25525511135"))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(ThreeSum([]int{0, 0, 0}))
}
