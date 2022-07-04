package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func newNode(v interface{}) *Node {
	return &Node{value: v, next: nil}
}

func newQueue() *Queue {
	return &Queue{first: nil, last: nil, length: 0}
}

func peek(q *Queue) {
	if q.length == 0 {
		fmt.Println("QUEUE is EMPTY")
	} else {
		fmt.Println(q.first.value, "\t len:", q.length)
	}
}
func enqueue(v interface{}, q *Queue) {
	n := newNode(v)
	if q.length == 0 {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		q.last = n
	}
	q.length++
}

func dequeue(q *Queue) *Node {
	if q.length > 0 {
		n := q.first
		q.first = q.first.next
		q.length--
		return n
	}
	return nil
}

func main() {
	q := newQueue()
	enqueue("google", q)
	enqueue("netflix", q)
	peek(q)
	dequeue(q)
	peek(q)
	dequeue(q)
	peek(q)
}
