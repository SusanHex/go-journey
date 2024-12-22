package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		data: data,
		next: nil,
	}
	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
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
		fmt.Println("Index:", index, "Current node:", &curr, "Data:", curr.data, "Next node:", curr.next)
		curr = curr.next
		index++
	}
}
