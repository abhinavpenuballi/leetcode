package main

type stack []int

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))

	stack := stack{}

	for i, temperature := range temperatures {
		for len(stack) > 0 && temperatures[stack.top()] < temperature {
			ans[stack.top()] = i - stack.pop()
		}

		stack.push(i)
	}

	return ans
}

func (stack *stack) push(val int) {
	*stack = append(*stack, val)
}

func (stack *stack) pop() int {
	last := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1]

	return last
}

func (stack *stack) top() int {
	return (*stack)[len(*stack)-1]
}
