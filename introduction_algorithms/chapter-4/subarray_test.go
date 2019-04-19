package chapter_4

import (
	"fmt"
	"testing"
)

func TestFindMaximumSubArray(t *testing.T) {
	fmt.Println(FindMaximumSubArray([]int{1, -1, 2, 3, 4}, 0, 4))
}
