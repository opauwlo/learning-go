package main

import (
  "fmt"
)

type True int

var verdade True

func main() {
  fmt.Println("\nExercicio 4, tipos personalizados and more")
  fmt.Println("==========================================\n")
  fmt.Printf("%v \t%T\n", verdade, verdade)

  verdade = 42

  fmt.Printf("%v\n", verdade)

}
