package insert

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	fmt.Println(SortInsert([]int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3, 1, 2, 3}, true))
}
