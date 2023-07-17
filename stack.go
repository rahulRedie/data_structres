package main

import (
	"fmt"
	"sync"
)

type stack struct {
	data []int
	lock sync.RWMutex
}

func (s *stack) push(element int) {
	defer s.lock.Unlock()
	s.lock.Lock()
	s.data = append(s.data, element)
}

func (s *stack) top() (int, error) {
	defer s.lock.Unlock()
	if len(s.data) > 0 {
		s.lock.Lock()
		return s.data[len(s.data)-1], nil
	}
	return 0, fmt.Errorf("stack is empty")
}

func (s *stack) pop() error {
	defer s.lock.Unlock()
	if len(s.data) > 0 {
		s.lock.Lock()
		s.data = s.data[:len(s.data)-1]
		return nil
	}
	return fmt.Errorf("stack is empty")
}

func (s *stack) size() int {
	return len(s.data)
}

func main() {
	stack := &stack{
		data: make([]int, 0),
	}
	fmt.Println("size of stack", stack.size())
	stack.push(12)
	stack.push(13)
	stack.push(15)
	for len(stack.data) > 0 {
		topValue, _ := stack.top()
		fmt.Println("Popping top value", topValue)
		stack.pop()
	}
	fmt.Println("final size", stack.size())
}
