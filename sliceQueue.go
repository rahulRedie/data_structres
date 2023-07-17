package main

import (
	"errors"
	"fmt"
	"sync"
)

type queue struct {
	slice []int
	lock  sync.RWMutex
}

func (c *queue) size() int {
	defer c.lock.Unlock()
	c.lock.Lock()
	return len(c.slice)
}

func (c *queue) front() (int, error) {
	defer c.lock.Unlock()
	if len(c.slice) > 0 {
		c.lock.Lock()
		return c.slice[0], nil
	}
	return 0, errors.New("length of slcie is zero")
}

func (c *queue) Enqueue(element int) {
	defer c.lock.Unlock()
	c.lock.Lock()
	c.slice = append(c.slice, element)
}

func (c *queue) Deque() error {
	defer c.lock.Unlock()
	if len(c.slice) > 0 {
		c.lock.Lock()
		c.slice = c.slice[1:]
		return nil
	}
	return fmt.Errorf("length is zero")
}

func main() {
	queue := &queue{
		slice: make([]int, 0),
	}
	fmt.Println("Initial length", queue.size())
	queue.Enqueue(12)
	queue.Enqueue(13)
	queue.Enqueue(15)

	for len(queue.slice) > 0 {
		queueelement, _ := queue.front()
		fmt.Println("dequeing front Value", queueelement)
		queue.Deque()
	}
	fmt.Println("final size", queue.size())
}
