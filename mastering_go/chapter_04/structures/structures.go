package main

import "fmt"

type node[T comparable] struct {
	data     T
	next     *node[T]
	previous *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) delete(inputNode *node[T]) {

	current_node := l.start
	for {
		if current_node == nil {
			return
		}
		if current_node == inputNode {
			break
		}
		current_node = current_node.next
	}
	previous_node := current_node.previous
	next_node := current_node.next
	if previous_node != nil {
		previous_node.next = next_node
	}
	if next_node != nil {
		next_node.previous = next_node
	}
	current_node.next = nil
	current_node.previous = nil
}

func (l *list[T]) print() {
	current_node := l.start
	for {
		if current_node == nil {
			break
		}
		fmt.Print(current_node.data, " ")
		current_node = current_node.next
	}
	fmt.Println()
}

func (l *list[T]) search(term T) *node[T] {
	current_node := l.start
	for {
		if current_node == nil {
			return nil
		}
		if current_node.data == term {
			return current_node
		}
		current_node = current_node.next
	}
}

func (l *list[T]) add(data T) {
	n := node[T]{
		data: data,
		next: nil,
	}
	if l.start == nil {
		l.start = &n
		l.start.previous = nil
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		l.start.next.previous = l.start
		return
	}
	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(1)
	myList.add(2)
	myList.add(3)
	myList.add(4)
	myList.add(5)
	fmt.Println(myList)
	curr := myList.start
	index := 0
	for {
		if curr == nil {
			break
		}
		fmt.Println("Index:", index, "Current node:", curr, "Data:", curr.data, "Next node:", curr.next)
		curr = curr.next
		index++
	}
	myList.print()
	tempNode := myList.search(2)
	if tempNode != nil {
		fmt.Println("Search list for 2, found node at:", tempNode)
		myList.delete(tempNode)
		fmt.Println("Deleted found node")
		myList.print()
	}
}
