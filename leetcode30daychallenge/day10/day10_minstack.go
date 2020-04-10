package day10

// MinStack is a stack that can report the minimum value in the stack
type MinStack struct {
	values []int
}

// Constructor creates a MinStack
func Constructor() MinStack {
	return MinStack{
		values: make([]int, 0),
	}
}

// Push a value onto the stack
func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
}

// Pop removes the top element from the stack
func (this *MinStack) Pop() {
	l := len(this.values)
	if l <= 0 {
		return
	}

	this.values = this.values[0 : l-1]
}

// Top returns the top element from the stack
func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

// GetMin finds the minimum value in the stack
func (this *MinStack) GetMin() int {
	min := this.values[0]

	for _, val := range this.values {
		if val < min {
			min = val
		}
	}

	return min
}
