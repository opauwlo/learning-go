package main

import (
  "fmt"
)


/* 
 Crie um programa que:
  Atribua um valor int a uma variável
  Demonstre este valor em decimal, binário e hexadecimal
  Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
  Demonstre esta outra variável em decimal, binário e hexadecimal
*/

var num int

func main() {
  fmt.Println("Exercicio 04")
  num = 2

  fmt.Printf("decimal: %v\t binário: %b\t hexadecimal: %#x\n", num, num, num)

  y := num << 1 // multiplica por 2

  fmt.Printf("decimal: %v\t binário: %b\t hexadecimal: %#x\n", y, y, y)

}

