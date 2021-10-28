package main

import "fmt"

/*
 * - Base-10: decimal, 0-9
 * - Base-2: binário, 0-1
 * - Base-16: hexadecimal, 0-f
 */

func show_bases() {
  var exit_message string
  var num int
  for true {
    fmt.Print("\n=============================================================================\n")
    fmt.Print("Insira o valor em decimal que deseja converter para binario e hexadecimal: ")
    fmt.Scanf("%d", &num)
    fmt.Printf("\nbinário: %b\t hexadecimal: %#x\n\n", num, num) // %d == decimal, %b == binário & %#x == hexadecimal
    fmt.Print("Digite exit para sair ou qualquer tecla para continuar: ")
    fmt.Scanf("%s",&exit_message)
    if (exit_message == "exit") {
      break
    } else {
      continue
    }
  }
 }

func main() {
  fmt.Println("Aula 08, Sistemas númericos")
  show_bases()
}
