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

func TestMaxAreaOfIsland(t *testing.T) {
	fmt.Println(MaxAreaOfIsland([][]int{
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}

func TestSearch(t *testing.T) {
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 10,
				},
			},
		},
	}
	fmt.Println(LevelOrder(root))
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val: 15,
			},
			Left: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(LevelOrder(root1))
}

func TestZigzagLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(ZigzagLevelOrder(root))
}
