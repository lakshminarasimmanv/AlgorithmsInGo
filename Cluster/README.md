 # Sequence Clustering Algorithm

 Go package for Sequence Clustering Algorithm.

 ## Installation

 ```
 go get github.com/lakshminarasimmanv/GoLearn/sca
 ```

 ## Usage

 ```go
 package main

 import (
 	"fmt"

 	"github.com/lakshminarasimmanv/GoLearn/sca"
 )

 func main() {
 	sequences := sca.GenerateSequences(10, 5, 0, 10)
 	sequences.Sort()

 	sca := sca.NewSCA(sequences, 3, 100, 0.001)
 	sca.Run()
 	sca.Print()
 }

## How it works?

- The program generates a slice of sequences.
- The sequences are sorted.
- The program runs the Sequence Clustering Algorithm.
- The program prints the clusters.