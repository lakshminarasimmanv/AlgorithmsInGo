/* Breadth-first Search Algorithm using Go. */

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if n.value == value {
		return
	} else if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) print() {
	if n == nil {
		return
	}
	n.left.print()
	fmt.Println(n.value)
	n.right.print()
}

func (n *Node) breadthFirstSearch(value int) *Node {
	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.value == value {
			return node
		}
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		return
	}
	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a number")
		return
	}
	root := &Node{value: 10}
	root.insert(5)
	root.insert(15)
	root.insert(20)
	root.insert(0)
	root.insert(-5)
	root.insert(3)
	root.print()
	fmt.Println("Searching for", value)
	node := root.breadthFirstSearch(value)
	if node != nil {
		fmt.Println("Found", value)
	} else {
		fmt.Println("Not found")
	}
}
