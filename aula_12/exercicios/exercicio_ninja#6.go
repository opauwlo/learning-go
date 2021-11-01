package main

import (
  "fmt"
)

/*
 * - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
 * - Demonstre estes valores.
 */

const (
  _ = 2021 +  iota
  ano22
  ano23
  ano24
  ano25
)

func main() {
  fmt.Println(ano22, ano23, ano24, ano25)
}
