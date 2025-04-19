//leetcode Q232 Implement Queue using Stacks

package classtwo

type MyQueue struct {
	inStack  []int //1,2
	outStack []int //2,1  pop inStack會得到反序的stack 就能實現FIFO 題目有點難懂意思
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	//先把inStack倒進outStack
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	//1,2 => 2,1 outStack's Top is first in
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}

func (this *MyQueue) Peek() int {
	//先把inStack倒進outStack
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

//Runtime 0ms Beats 100.00%
//Memory 4.12MB Beats 6.34%
