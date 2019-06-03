package quick

import (
	"fmt"
	"testing"
)

func TestSortQuick(t *testing.T) {
	fmt.Println(SortQuick([]int{1, 3, 2, 2, 4}))
	fmt.Println(SortQuick2([]int{1, 2, 3, 3, 5, 4}))
}
