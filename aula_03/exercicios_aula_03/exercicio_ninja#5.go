package main

import (
  "fmt"
)

type True int

var(
  verdade True
  y int
)
func main() {
  fmt.Println("\nExercicio 5, tipos personalizados e conversion")
  fmt.Println("==========================================")
  fmt.Printf("%v \t%T\n", verdade, verdade)
  fmt.Printf("%v \t%T\n", y, y)

  verdade = 42
  y = int(verdade)
  fmt.Printf("%v \t%T\n", verdade, verdade)
  fmt.Printf("%v \t%T\n", y, y)

}
