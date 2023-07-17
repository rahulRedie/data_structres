package main

import (
	"container/list"
	"errors"
	"fmt"
)

type queue struct {
	data   *list.List
	length int
}

func (c *queue) front() (int, error) {
	if c.data.Len() > 0 {
		if val, ok := c.data.Front().Value.(int); ok {
			return val, nil
		}
		return 0, errors.New("error in queue")
	}
	return 0, fmt.Errorf("queue is empty")
}

func (c *queue) size() int {
	return c.data.Len()
}

func (c *queue) Enqueue(qelement int) error {
	c.data.PushBack(qelement)
	return nil
}

func (c *queue) Dequeue() error {
	if c.data.Len() > 0 {
		ele := c.data.Front()
		c.data.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}
func (c *queue) empty() bool {
	return c.data.Len() == 0
}

// func main() {
// 	queue := &queue{
// 		data: list.New(),
// 	}
// 	queue.Enqueue(12)
// 	fmt.Println(queue)
// 	fmt.Println("length", queue.size())
// 	ele, _ := queue.front()
// 	fmt.Println("front", ele)
// }
