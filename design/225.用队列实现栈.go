package design

//  description: (https://leetcode.cn/problems/implement-stack-using-queues/)

//  你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
//	实现 MyStack 类：
//	void push(int x) 将元素 x 压入栈顶。
//	int pop() 移除并返回栈顶元素。
//	int top() 返回栈顶元素。
//	boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

type MyStack struct {
	queue []int
}

func NewMyStack() MyStack {
	stk := MyStack{}
	stk.queue = []int{}
	return stk
}

func (s *MyStack) Push(x int) {
	l := len(s.queue)
	s.queue = append(s.queue, x)
	for i := 0; i < l; i++ {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *MyStack) Pop() int {
	top := s.queue[0]
	s.queue = s.queue[1:]
	return top
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
