//linked list - addfront, addback, removefront, removeback, front, size, traverse, initlist

package main

import (
	"errors"
	"fmt"
)

type element struct {
	data     int
	nextNode *element
}

type linkedlist struct {
	head *element
	len  int
}

func initlist() *linkedlist {
	return &linkedlist{}
}

func (l *linkedlist) AddFront(value int) {

	ele := &element{
		data: value,
	}
	if l.head == nil {
		l.head = ele
	} else {
		ele.nextNode = l.head
		l.head = ele
	}
	l.len++
	return
}

func (l *linkedlist) addback(value int) {
	ele := &element{
		data: value,
	}
	if l.head == nil {
		l.head = ele
	} else {
		current := l.head
		for current.nextNode != nil {
			current = current.nextNode
		}
		current.nextNode = ele
	}
	l.len++
	return
}

func (l *linkedlist) removefront() error {
	if l.head == nil {
		return errors.New("empty linked list")
	}
	l.head = l.head.nextNode
	l.len--
	return nil
}

func (l *linkedlist) removeback() error {
	if l.head == nil {
		return fmt.Errorf("empty linked list")
	}
	var previous *element
	current := l.head
	for current.nextNode != nil {
		previous = current
		current = current.nextNode
	}
	if previous != nil {
		previous.nextNode = nil
	} else {
		l.head = nil
	}
	current.nextNode = nil
	l.len--
	return nil
}
func traverse(l *linkedlist) {
	curent := l.head
	for curent != nil {
		fmt.Println(curent.data)
		curent = curent.nextNode
	}
}
func (l *linkedlist) size() (length int) {
	length = l.len
	return
}

func main() {
	list := initlist()
	list.AddFront(12)
	list.AddFront(19)
	list.addback(13)
	list.AddFront(14)
	fmt.Println(list.size())
}
