package main

type stack []int

func canSeePersonsCount(heights []int) []int {
	canSee := []int{0}

	stack := stack{heights[len(heights)-1]}

	for i := len(heights) - 2; i >= 0; i-- {
		thisPersonCanSee := 0

		for j := len(stack) - 1; j >= 0; j-- {
			thisPersonCanSee++

			if stack[j] >= heights[i] {
				break
			}
		}

		canSee = append([]int{thisPersonCanSee}, canSee...)

		for len(stack) > 0 && stack.top() <= heights[i] {
			stack.pop()
		}

		stack.push(heights[i])
	}

	return canSee
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
