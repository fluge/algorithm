package leetcode

type MinStack struct {
	val []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		val: make([]int, 0),
		min: 0,
	}
}

func (stack *MinStack) Push(x int) {

	if len(stack.val) == 0 {
		stack.min = x
	} else {
		if x < stack.min {
			stack.min = x
		}
	}
	stack.val = append(stack.val, x)

}

func (stack *MinStack) Pop() {

	//遍历整个栈找到最小的元素

	stack.val = stack.val[:len(stack.val)-1]

	if len(stack.val) == 0 {
		stack.min = MinInt
		return
	}
	stack.min = stack.val[0]
	for _, v := range stack.val {
		if v < stack.min {
			stack.min = v
		}
	}
}

func (stack *MinStack) Top() int {

	return stack.val[len(stack.val)-1]
}

func (stack *MinStack) GetMin() int {

	return stack.min
}
