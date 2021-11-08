package main 

import (
  "fmt"
)

func main() {
  for i := 33; i < 123; i++ {
    fmt.Printf("decimal: %v\tASCII: %v\n", i, string(i))
  }

}
