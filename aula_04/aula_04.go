package main

import (
  "fmt"
)

var x bool

func main() {
  fmt.Println("Aula 04, tipo booleano")
  fmt.Printf("\n%v\n", x) // zero value
  x = true
  fmt.Println(x) //valor atribuido
  // operadores relacionais
  x = (10 < 100)
  y := (10 == 100)
  z := (10 > 100)


  fmt.Println(x) //valor atribuido

  fmt.Println(y) //valor atribuido

  fmt.Println(z) //valor atribuido

}
