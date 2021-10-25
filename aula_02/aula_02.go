package main

import (
  "fmt"
)

func func_with_params(age int, name string, live bool) {
  fmt.Println(age, name, live)
}

func for_star_ex() {
  for i:=1; i<=5; i++ {
    for j:=5; j>=i; j-- {
      fmt.Print("*")
    }
    fmt.Print("\n")
  }
}

func for_star2_ex() {
  for i:=1; i<=5; i++ {
      for j:=1; j<=i; j++ {
      fmt.Print("*")
    }
    fmt.Print("\n")
  }
}

func main() {
  fmt.Println("aula 02\nrec\n")
  // func_with_params(21, "pauwlo", true)
  for_star_ex()
  for_star2_ex()
}
