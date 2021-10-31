package main

import (
  "fmt"
)

// escreva um programa que mostre um número decimal, binário e hexadecimal.

func main() {
  num := 10

  fmt.Printf("%v, %b, %#x\n", num, num, num)
}
