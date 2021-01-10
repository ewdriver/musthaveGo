//ex22.3/ex22.3.go
package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.l.PushBack(val) // ❶ 맨 뒤에 요소 추가
}

func (s *Stack) Pop() interface{} {
	back := s.l.Back() // ❷ 맨 뒤에서 요소를 반환
	if back != nil {
		return s.l.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
