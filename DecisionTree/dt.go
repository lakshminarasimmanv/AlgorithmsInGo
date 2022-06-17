/* Go package for Decision Tree. */
package dt

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// DecisionTree is a struct for a decision tree.
type DecisionTree struct {
	root *Node
}

// Node is a struct for a node in a decision tree.
type Node struct {
	attribute int
	value     float64
	left      *Node
	right     *Node
	label     int
}

// NewDecisionTree returns a new decision tree.
func NewDecisionTree() *DecisionTree {
	return &DecisionTree{}
}

// Train trains a decision tree.
func (dt *DecisionTree) Train(data [][]float64, labels []int, attributes []int) {
	dt.root = dt.train(data, labels, attributes)
}

// train trains a decision tree.
func (dt *DecisionTree) train(data [][]float64, labels []int, attributes []int) *Node {
	if len(labels) == 0 {
		return nil
	}
	if len(attributes) == 0 {
		return &Node{label: majority(labels)}
	}
	if isSame(labels) {
		return &Node{label: labels[0]}
	}
	best := bestAttribute(data, labels, attributes)
	if best == -1 {
		return &Node{label: majority(labels)}
	}
	leftData, leftLabels, rightData, rightLabels := split(data, labels, best)
	leftAttributes := remove(attributes, best)
	rightAttributes := remove(attributes, best)
	return &Node{
		attribute: best,
		value:     median(data, best),
		left:      dt.train(leftData, leftLabels, leftAttributes),
		right:     dt.train(rightData, rightLabels, rightAttributes),
	}
}

// Predict predicts the label of a data point.
func (dt *DecisionTree) Predict(data []float64) int {
	return dt.predict(dt.root, data)
}

// predict predicts the label of a data point.
func (dt *DecisionTree) predict(node *Node, data []float64) int {
	if node.left == nil && node.right == nil {
		return node.label
	}
	if data[node.attribute] < node.value {
		return dt.predict(node.left, data)
	}
	return dt.predict(node.right, data)
}

// Print prints the decision tree.
func (dt *DecisionTree) Print() {
	dt.print(dt.root, 0)
}

// print prints the decision tree.
func (dt *DecisionTree) print(node *Node, depth int) {
	if node == nil {
		return
	}
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	if node.left == nil && node.right == nil {
		fmt.Println(node.label)
		return
	}
	fmt.Printf("%d < %.2f\n", node.attribute, node.value)
	dt.print(node.left, depth+1)
	dt.print(node.right, depth+1)
}

// bestAttribute returns the best attribute to split on.
func bestAttribute(data [][]float64, labels []int, attributes []int) int {
	best := -1
	bestGain := -1.0
	for _, attribute := range attributes {
		gain := informationGain(data, labels, attribute)
		if gain > bestGain {
			best = attribute
			bestGain = gain
		}
	}
	return best
}

// informationGain returns the information gain of an attribute.
func informationGain(data [][]float64, labels []int, attribute int) float64 {
	leftData, leftLabels, rightData, rightLabels := split(data, labels, attribute)
	return entropy(labels) -
		(float64(len(leftLabels))/float64(len(labels))*entropy(leftLabels) +
			float64(len(rightLabels))/float64(len(labels))*entropy(rightLabels))
}

// entropy returns the entropy of a set of labels.
func entropy(labels []int) float64 {
	counts := make(map[int]int)
	for _, label := range labels {
		counts[label]++
	}
	entropy := 0.0
	for _, count := range counts {
		p := float64(count) / float64(len(labels))
		entropy -= p * math.Log2(p)
	}
	return entropy
}

// split splits a dataset into two parts.
func split(data [][]float64, labels []int, attribute int) ([][]float64, []int, [][]float64, []int) {
	leftData := make([][]float64, 0)
	leftLabels := make([]int, 0)
	rightData := make([][]float64, 0)
	rightLabels := make([]int, 0)
	for i, datum := range data {
		if datum[attribute] < median(data, attribute) {
			leftData = append(leftData, datum)
			leftLabels = append(leftLabels, labels[i])
		} else {
			rightData = append(rightData, datum)
			rightLabels = append(rightLabels, labels[i])
		}
	}
	return leftData, leftLabels, rightData, rightLabels
}

// median returns the median of an attribute.
func median(data [][]float64, attribute int) float64 {
	values := make([]float64, len(data))
	for i, datum := range data {
		values[i] = datum[attribute]
	}
	sort.Float64s(values)
	return values[len(values)/2]
}

// majority returns the majority label.
func majority(labels []int) int {
	counts := make(map[int]int)
	for _, label := range labels {
		counts[label]++
	}
	max := -1
	maxCount := -1
	for label, count := range counts {
		if count > maxCount {
			max = label
			maxCount = count
		}
	}
	return max
}

// isSame returns true if all labels are the same.
func isSame(labels []int) bool {
	for i := 1; i < len(labels); i++ {
		if labels[i] != labels[0] {
			return false
		}
	}
	return true
}

// remove removes an element from a slice.
func remove(slice []int, element int) []int {
	for i, e := range slice {
		if e == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// RandomForest is a struct for a random forest.
type RandomForest struct {
	trees []*DecisionTree
}

// NewRandomForest returns a new random forest.
func NewRandomForest(n int) *RandomForest {
	return &RandomForest{
		trees: make([]*DecisionTree, n),
	}
}

// Train trains a random forest.
func (rf *RandomForest) Train(data [][]float64, labels []int, attributes []int) {
	for i := range rf.trees {
		rf.trees[i] = NewDecisionTree()
		rf.trees[i].Train(data, labels, attributes)
	}
}

// Predict predicts the label of a data point.
func (rf *RandomForest) Predict(data []float64) int {
	counts := make(map[int]int)
	for _, tree := range rf.trees {
		label := tree.Predict(data)
		counts[label]++
	}
	max := -1
	maxCount := -1
	for label, count := range counts {
		if count > maxCount {
			max = label
			maxCount = count
		}
	}
	return max
}

// Print prints the random forest.
func (rf *RandomForest) Print() {
	for i, tree := range rf.trees {
		fmt.Printf("Tree %d:\n", i)
		tree.Print()
	}
}

// RandomForestBagging is a struct for a random forest using bagging.
type RandomForestBagging struct {
	trees []*DecisionTree
}

// NewRandomForestBagging returns a new random forest using bagging.
func NewRandomForestBagging(n int) *RandomForestBagging {
	return &RandomForestBagging{
		trees: make([]*DecisionTree, n),
	}
}

// Train trains a random forest using bagging.
func (rf *RandomForestBagging) Train(data [][]float64, labels []int, attributes []int) {
	for i := range rf.trees {
		rf.trees[i] = NewDecisionTree()
		rf.trees[i].Train(bagging(data, labels), labels, attributes)
	}
}

// Predict predicts the label of a data point.
func (rf *RandomForestBagging) Predict(data []float64) int {
	counts := make(map[int]int)
	for _, tree := range rf.trees {
		label := tree.Predict(data)
		counts[label]++
	}
	max := -1
	maxCount := -1
	for label, count := range counts {
		if count > maxCount {
			max = label
			maxCount = count
		}
	}
	return max
}

// Print prints the random forest using bagging.
func (rf *RandomForestBagging) Print() {
	for i, tree := range rf.trees {
		fmt.Printf("Tree %d:\n", i)
		tree.Print()
	}
}

// bagging returns a bagged dataset.
func bagging(data [][]float64, labels []int) [][]float64 {
	baggedData := make([][]float64, len(data))
	rand.Seed(time.Now().UnixNano())
	for i := range baggedData {
		baggedData[i] = data[rand.Intn(len(data))]
	}
	return baggedData
}
