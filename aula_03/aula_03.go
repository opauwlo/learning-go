package main

import (
  "fmt"
)

type age int

var idade age = 21

func main() {
  fmt.Println("aula 03, tipos são vida")

  fmt.Println(idade)

  // convertendo do tipo age para int
  idade_int := 21

  // soma das idades, preciso primeiro passar os tipos
  // aqui tranformamos a var idadee_int para o tipo age
  soma_das_idades := idade + age(idade_int)
  // notação simples tipo(bar)

  fmt.Println(soma_das_idades)
}
