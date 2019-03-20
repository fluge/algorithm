package merging

import (
	"fmt"
	"testing"
)

func TestSortMerge(t *testing.T) {
	fmt.Println(SortMerge([]int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 1, 2, 3}))
}
