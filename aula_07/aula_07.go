package main

import (
  "fmt"
)

/*
 * - Strings são sequencias de bytes.
 * - Imutáveis
 * - Uma String é um "slice of bytes" (ou, em português, uma faatia de bytes).
 * - Na prática:
 *   - %v %T
 *   - Raw string literals
 *   - Conversão par sslice of bytes []byte(x)
 *   - %#U, %#x
 * 
 */


var s string = "blá blá, ĸß?→/ß→→//"

func for_with_range() {
  s := "blá blá, ĸß?→/ß→→//"
  
  for _, v := range s { // passa pela string char por char
    fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
  }
  fmt.Print("\n")
}

func for_tradicional() {
  for i := 0; i < len(s); i++ { // passa pela string byte por byte
    fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
  }
}

func main() {
  intro := "Aula 07, tipo string\n" // interpreted string the \n will be interpreted an break one line
  intro_raw := `Aula 07 
                        tipo:
                              string\n` // raw string the \n will be not interpreted
  fmt.Println(intro)
  fmt.Printf("tipo: %T\n", intro) // tipo dda variavel intro == string
  fmt.Println(intro_raw)
  fmt.Printf("tipo: %T\n", intro_raw) // tipo dda variavel intro == string
  slice_intro := []byte(intro) // qual byte representa cada char
  fmt.Println(slice_intro) // mostra na tela
  for_with_range()
  for_tradicional()
}
