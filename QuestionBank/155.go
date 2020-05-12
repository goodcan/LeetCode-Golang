/*
 * @Time     : 2020/5/12 9:36
 * @Author   : cancan
 * @File     : 155.go
 * @Function : 最小栈
 */

/*
 * Question:
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 * 1.push(x) -- 将元素 x 推入栈中。
 * 2.pop() -- 删除栈顶的元素。
 * 3.top() -- 获取栈顶元素。
 * g4.etMin() -- 检索栈中的最小元素。
 *
 * Example:
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 */

package QuestionBank

type MinStack struct {
	d [][2]int
	l int
}

/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	m := x
	if this.l > 0 && m > this.d[this.l-1][1] {
		m = this.d[this.l-1][1]
	}
	this.d = append(this.d, [2]int{x, m})
	this.l++
}

func (this *MinStack) Pop() {
	this.d = this.d[:this.l-1]
	this.l--
}

func (this *MinStack) Top() int {
	return this.d[this.l-1][0]
}

func (this *MinStack) GetMin() int {
	return this.d[this.l-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
