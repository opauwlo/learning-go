package main

import (
	"fmt"
	"math"
)

const (
  _ = iota * math.Pi // numa declaração de consts, o identificador iota representa numeros sequenciais.
  x
  y
  z
)
func main() {
  fmt.Println("Aula 10, Iota")
  fmt.Printf("%T\t -\t%v\n%T\t -\t%v\n%T\t -\t%v\n", x, x, y, y, z, z)
}
