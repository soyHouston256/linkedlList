package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	fmt.Println("Linked List")
	myList := LinkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}

	myList.prepend(node3)
	myList.prepend(node2)
	myList.prepend(node1)
	fmt.Println(myList)
}
