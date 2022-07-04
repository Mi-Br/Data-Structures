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

	head := l.head
	tail := l.tail

	first := l.head
	second := l.head.next
	fmt.Println("doing smth")
	//swap next reference
	for second != nil {
		tmp := second.next
		second.next = first
		first, second = second, tmp
	}
	//swap head and tail
	tmp := head
	head = tail
	tail = tmp
	fmt.Println("doing smth")
}

func main() {
	var list = new(LinkedList)
	// fmt.Println(list)
	appendList(10, list)
	appendList(5, list)
	appendList(1, list)
	printNodes(list)
	reverseList(list)
	fmt.Println("head:  with value of: \t ", &list.head, list.head.value)
	fmt.Println("tail:  with value of: \t ", &list.tail, list.tail.value)
	fmt.Println(list.tail.next)
	fmt.Println(list.head.next)

	// printNodes(list)
	// appendList(1, list)
	// appendList(3, list)
	// printNodes(list)
	// prependList(-3, list)
	// prependList(-2, list)
	// prependList(-1, list)
	// printNodes(list)
	// insertNode(100, 0, list)
	// printNodes(list)
	// insertNode(90, 10, list)
	// printNodes(list)
	// insertNode(1000, 5, list)
}
