// Queue using linked list in Go.
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(data int) {
	newNode := &Node{data, nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		fmt.Println("Queue is empty.")
		return -1
	}
	data := q.head.data
	q.head = q.head.next
	return data
}

func (q *Queue) Print() {
	if q.head == nil {
		fmt.Println("Queue is empty.")
		return
	}
	for curr := q.head; curr != nil; curr = curr.next {
		fmt.Printf("%d ", curr.data)
	}
	fmt.Println()
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Print()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Print()
}
