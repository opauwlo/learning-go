package main

import (
  "fmt"
)

func main() {
  fmt.Println("Aula 15, loops: a declaração")
  
  i := 0

  for true {
    fmt.Println("o for é true")
    i++

    if (i >= 10) {
      fmt.Println("o for é false mano, corre")
      break // para o for eterno ou qualquer for
    } 
  }

  fmt.Println("ufa")
}
