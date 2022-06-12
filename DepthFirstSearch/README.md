# Depth First Search

The code above implements a depth-first search algorithm in Go. A depth first search is a type of search algorithm that traverses a tree or graph by first visiting the nodes at the deepest level, and then working its way back up to the root.

To use the depth first search algorithm, simply create a new Node, and insert values into it using the insert() method. Then, call the depthFirstSearch() method on the root node, and pass in an empty slice of ints. This will return a slice of ints containing all the values in the tree or graph, in depth first order.

## Installation

You will need to have the Go programming language installed on your system.

## Usage

```
go run depth-first-search.go [list of integers to sort]
```

## Example

```
go run depth-first-search.go 1 2 3 4 5
```

This will output the following:

```
[1 2 3 4 5]
```
