/* Depth-First Search Algorithm using Go. */

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

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n *Node) insert(value int) {
	if value <= n.value {
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

func (n *Node) contains(value int) bool {
	if n.value == value {
		return true
	} else if value < n.value {
		if n.left == nil {
			return false
		} else {
			return n.left.contains(value)
		}
	} else {
		if n.right == nil {
			return false
		} else {
			return n.right.contains(value)
		}
	}
}

func (n *Node) depthFirstSearch(values []int) []int {
	values = append(values, n.value)
	if n.left != nil {
		values = n.left.depthFirstSearch(values)
	}
	if n.right != nil {
		values = n.right.depthFirstSearch(values)
	}
	return values
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a list of integers to sort.")
		os.Exit(1)
	}

	var root *Node
	for _, arg := range os.Args[1:] {
		value, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%q is not an integer.\n", arg)
			os.Exit(1)
		}
		if root == nil {
			root = &Node{value: value}
		} else {
			root.insert(value)
		}
	}

	fmt.Println(root.depthFirstSearch([]int{}))
}
