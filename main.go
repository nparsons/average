// average calculates the average of several numbers in a set.
package main

import (
  "fmt"
  "log"
  "github.com/nparsons/datafile"
  "os"
  "strconv"
)

func main() {
  var numbers []float64
  var err error

  if len(os.Args) > 1 {
    for _, arg := range os.Args[1:] {
      num, err := strconv.ParseFloat(arg, 64)
      if err != nil {
        log.Fatal(err)
      }
      numbers = append(numbers, num)
    }
  } else {
    numbers, err = datafile.GetFloats("data.txt")

    if err != nil {
      log.Fatal(err)
    }
  }

  var sum float64 = 0
  for _, number := range numbers {
    sum += number
  }

  sampleCount := float64(len(numbers))
  fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}
