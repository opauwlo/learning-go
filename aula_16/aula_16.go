package main

import (
  "fmt"
)

func main() {
  fmt.Println("Aula, loops: continue")

  for i := 0; i <= 5; i++ {
    if (i%2 != 0) {
      fmt.Printf("%v é impar\n", i)
      continue // quebra o loop no meio e volta pro começo dele
    }
    fmt.Printf("%v é par\n", i)
  }
}
