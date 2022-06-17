# Go package for Time Series

This Go package provides a time series type.

## Installation

```
go get github.com/lakshminarasimmanv/GoLearn/go-ts
```

## Example

Here is an example program.

```
     package main

     import (
         "fmt"
         "time"

         "github.com/soniakeys/ts"
     )

     func main() {
         ts := ts.NewTimeSeries("foo", map[string]string{"bar": "baz"})
         ts.AddPoint(time.Now(), 1)
         ts.AddPoint(time.Now(), 2)
         ts.AddPoint(time.Now(), 3)
         ts.Sort()
         fmt.Println(ts)
         ts = ts.Aggregate("sum", time.Hour)
         fmt.Println(ts)
         ts, _ = ts.ParseTimeSeries("foo{bar=\"baz\"} 1546300800000 6 1546301700000 9")
         fmt.Println(ts)
         ts = ts.Aggregate("sum", time.Hour)
         fmt.Println(ts)
     }
```

Here is the output from the example program.

```
     foo{bar="baz"} 1546300800000 1 1546301700000 2 1546302600000 3
     foo{bar="baz"} 1546300800000 6
     foo{bar="baz"} 1546300800000 6 1546301700000 9
     foo{bar="baz"} 1546300800000 15
```

## Documentation

 Documentation is available at godoc.org.

## Tests

To run the tests, run

     `go test`

## Benchmarks

To run the benchmarks, run

     `go test -bench .`

## License

This Go package is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.