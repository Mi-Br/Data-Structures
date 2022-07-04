package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewNode(v int) *Node {
	return &Node{value: v, next: nil}
}

func appendList(v int, l *LinkedList) *LinkedList {
	n := NewNode(v)
	if l.head == nil {
		l.head = n
	}
	if l.tail == nil {
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.length++
	return l
}

func prependList(v int, l *LinkedList) *LinkedList {
	n := NewNode(v)
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.length++
	return l
}

func printNodes(l *LinkedList) {
	p := l.head
	fmt.Print("head|")
	for p != nil {
		fmt.Print("[", p.value, "]->")
		p = p.next
	}
	fmt.Println(" |tail")
	fmt.Println("list length: ", l.length)
}

func insertNode(v int, s int, l *LinkedList) *LinkedList {
	n := NewNode(v)
	p := l.head
	if p == nil || s >= l.length { //if insert used on an empty list - just use append
		appendList(v, l)
	} else if s <= 0 {
		prependList(v, l)
	} else {
		for p != nil && s != 0 {
			p = p.next
			s = s - 1
		}
		n.next = p.next
		p.next = n
		l.length++
	}

	return l
}

func reverseList(l *LinkedList) {
	fmt.Println(l.length)
	if l.length < 2 {
		fmt.Println("list too short, returning it")
		return
	}

	first := l.head
	second := l.head.next

	//swap next reference
	for second != nil {
		tmp := second.next
		second.next = first
		first, second = second, tmp
	}
	//swap head and tail
	l.head, l.tail = l.tail, l.head
	l.tail.next = nil
}

func main() {
	var list = new(LinkedList)
	// fmt.Println(list)
	appendList(10, list)
	appendList(5, list)
	appendList(1, list)
	printNodes(list)
	reverseList(list)
	printNodes(list)
}
