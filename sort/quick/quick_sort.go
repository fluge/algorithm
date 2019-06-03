package quick

/*
快排是不稳定排序，平均是O(n*log n) ，在整个逆序的情况下时间复杂度会退化为O(n^2)
*/
//经典实现
func SortQuick(arr []int) []int {
	//就以左边的第一个为基准
	partition(arr, 0, len(arr)-1)
	return arr
}

func partition(arr []int, left, right int) {
	if right > left {
		key := arr[left]
		i, j := left, right
		for {
			for arr[j] >= key && i < j {
				j--
			}
			for arr[i] <= key && i < j {
				i++
			}
			if i >= j {
				break
			}

			arr[i], arr[j] = arr[j], arr[i]
		}

		arr[left] = arr[i]
		arr[i] = key
		partition(arr, left, i-1)
		partition(arr, i+1, right)
	}
}

//更加简洁的做法
func SortQuick2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	key, i := arr[0], 1
	head, tail := 0, len(arr)-1

	for head < tail {
		if arr[i] > key {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		} else {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	SortQuick2(arr[:head])
	SortQuick2(arr[head+1:])
	return arr
}
