package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.length++
}

func (l *linkedList) printListData() {
	toPrint := l.head
	length := l.length
	for length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	fmt.Println("Linked List")
	myList := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	node5 := &node{data: 5}
	node6 := &node{data: 6}
	node7 := &node{data: 7}
	node8 := &node{data: 8}
	node9 := &node{data: 9}
	node10 := &node{data: 10}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.append(node10)

	myList.printListData()

	myList.deleteWithValue(5)

	myList.printListData()
	myList.deleteWithValue(10)
	myList.printListData()
	myList.deleteWithValue(1)
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteWithValue(1)
}
