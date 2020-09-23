// average calculates the average of several numbers in a set.
package main

import (
  "fmt"
  "log"
  "github.com/nparsons/datafile"
)

func main() {
  numbers, err := datafile.GetFloats("data.txt")
  if err != nil {
    log.Fatal(err)
  }

  var sum float64 = 0
  for _, number := range numbers {
    sum += number
  }

  sampleCount := float64(len(numbers))
  fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}
