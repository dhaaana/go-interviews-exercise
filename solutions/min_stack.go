package solutions

import "fmt"

// Problem: https://leetcode.com/problems/min-stack/submissions/
// 155. Min Stack

type StackNode struct {
	node int
	min  int
}

type MinStack struct {
	nodes []*StackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	newStackNode := &StackNode{min: val, node: val}

	if len(this.nodes) > 0 {
		prevStackNode := this.nodes[len(this.nodes)-1]

		if prevStackNode.min < val {
			newStackNode.min = prevStackNode.min
		}
	}

	this.nodes = append(this.nodes, newStackNode)
}

func (this *MinStack) Pop() {
	this.nodes = this.nodes[:len(this.nodes)-1]
}

func (this *MinStack) Top() int {
	return this.nodes[len(this.nodes)-1].node
}

func (this *MinStack) GetMin() int {
	return this.nodes[len(this.nodes)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func MinStackSolution() {
	val := 1
	val2 := 2
	val3 := 0
	val4 := -3

	obj := Constructor()
	obj.Push(val)
	obj.Push(val2)
	obj.Push(val3)
	obj.Push(val4)

	fmt.Print("initial stack ")
	obj.PrintStack()

	obj.Pop()

	fmt.Print("after popped ")
	obj.PrintStack()

	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println("top element", param_3)
	fmt.Println("min element", param_4)
}

func (this *MinStack) PrintStack() {
	for _, i := range this.nodes {
		fmt.Print(i.node, " ")
	}
	fmt.Println()
}
