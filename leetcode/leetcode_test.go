package leetcode

import (
	"encoding/json"
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

func TestInvertTree(t *testing.T) {
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

	s := InvertTree2(root)
	str, _ := json.Marshal(s)
	fmt.Println(string(str))
}

func TestMaxDepth(t *testing.T) {
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
	fmt.Println(MaxDepth(root))
	fmt.Println(MaxDepth2(root))
}
func TestInorderTraversal(t *testing.T) {
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
	fmt.Println(InOrderTraversal(root))
	fmt.Println(InOrderTraversal2(root))

	fmt.Println(PreOrderTraversal(root))
	fmt.Println(PreOrderTraversal2(root))

	fmt.Println("后序：", PostOrderTraversal(root))
	fmt.Println("后序：", PostOrderTraversal2(root))

	fmt.Println(IsValidBST(root))
	fmt.Println(IsValidBST2(root))
}

func TestBuildTree(t *testing.T) {
	inO := []int{9, 3, 15, 20, 7}
	postO := []int{9, 15, 7, 20, 3}
	fmt.Println(BuildTree2(inO, postO))
}

func TestSortedListToBST(t *testing.T) {
	root := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 10,
					},
				},
			},
		},
	}
	fmt.Println(SortedListToBST(root))
}

func TestSearchMatrix(t *testing.T) {
	res := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(SearchMatrix(res, 5))
}

func TestFindLengthOfLCIS(t *testing.T) {
	//fmt.Println(FindLengthOfLCIS([]int{1, 3, 5, 4, 7, 8, 9, 10}))
	//fmt.Println(FindLengthOfLCIS([]int{2, 2, 2, 2}))
	//fmt.Println(FindLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5}))

	fmt.Println(LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	fmt.Println(LongestConsecutive([]int{-6, -1, -1, 9, -8, -6, -6, 4, 4, -3, -8, -1}))

	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))

}

func TestSubSets(t *testing.T) {
	fmt.Println(SubSets([]int{1, 2, 3}))
	fmt.Println(SubSets([]int{1, 2, 2}))

}

func TestAddTwoNumbers(t *testing.T) {
	node := addTwoNumbers(&ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 1,
		},
	}, &ListNode{
		Val: 9,
	})
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
}

func TestMergeTwoLists(t *testing.T) {
	node := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 7,
		},
	}, &ListNode{
		Val: 2,
	})
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
}

func TestRemoveNthFromEnd(t *testing.T) {
	node := removeNthFromEnd2(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, 5)
	str, _ := json.Marshal(node)
	fmt.Println(string(str))

	node = removeNthFromEnd(&ListNode{
		Val: 1,
	}, 1)
	str, _ = json.Marshal(node)
	fmt.Println(string(str))
}

func TestMergeSortList(t *testing.T) {
	node := mergeSortList(&ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	})
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
}

func TestMergeKLists(t *testing.T) {
	node := mergeKLists([]*ListNode{
		&ListNode{
			Val: 1,
		},
		&ListNode{
			Val: 2,
		},
		&ListNode{
			Val: 3,
		},
	})
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
}

func TestSwapPairs(t *testing.T) {
	node := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	})
	str, _ := json.Marshal(node)
	fmt.Println(string(str))
}

func TestRotateRight(t *testing.T) {
	node := rotateRight(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, 1)
	str, _ := json.Marshal(node)
	fmt.Println(string(str))

	node = rotateRights(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, 1)
	str, _ = json.Marshal(node)
	fmt.Println(string(str))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("ccc"))
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
