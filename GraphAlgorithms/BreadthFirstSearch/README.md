# Go Breadth-First Search

In this exercise, we'll implement a breadth-first search algorithm in Go.

## Creating the Node Struct

We'll start by creating a struct to represent a node in our tree. This struct will have three fields:

- `value`: the value stored at this node
- `left`: a pointer to the left child node
- `right`: a pointer to the right child node

Here's the code for our `Node` struct:

```go
type Node struct {
    value int
    left  *Node
    right *Node
}
```

## Implementing the `insert` Method

Next, we'll implement the `insert` method on our `Node` struct. This method will take a value as an argument, and add a new node containing that value to the tree.

If the value to insert is less than the value of the current node, it will be inserted to the left of the current node. If the value to insert is greater than the value of the current node, it will be inserted to the right.

Here's the code for the `insert` method:

```go
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
```

## Implementing the `print` Method

Next, we'll implement the `print` method on our `Node` struct. This method will print out the values of all the nodes in the tree, in ascending order.

Here's the code for the `print` method:

```go
func (n *Node) print() {
    if n == nil {
        return
    }
    n.left.print()
    fmt.Println(n.value)
    n.right.print()
}
```

## Implementing the `breadthFirstSearch` Method

Finally, we'll implement the `breadthFirstSearch` method on our `Node` struct. This method will take a value as an argument, and return the node in the tree containing that value.

The `breadthFirstSearch` method will use a queue to keep track of the nodes it needs to visit. It will start by adding the root node to the queue. Then, it will loop through the queue until it either finds the node it's looking for, or the queue is empty (meaning the node isn't in the tree).

For each node in the queue, the `breadthFirstSearch` method will first check to see if the node's value is the value it's looking for. If so, it will return that node. If not, it will add the node's left and right child nodes to the queue, and continue looping.

Here's the code for the `breadthFirstSearch` method:

```go
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
```

## Running the Code

To run the code, you can use the `go run` command:

```
go run main.go <value>
```

Where `<value>` is the value you want to search for in the tree.

For example, if you want to search for the value `5`, you would run the following command:

```
go run main.go 5
```

If the value is found in the tree, the program will print `Found <value>`. If the value is not found, the program will print `Not found`.
