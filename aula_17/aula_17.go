package main 

import (
  "fmt"
)

// Desafio supresa, tabela ascii

func main() {
  for i := 33; i < 123; i++ {
    fmt.Printf("Num: %v ASCII: %c\n", i, i)
  }

}
