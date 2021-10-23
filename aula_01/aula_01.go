package main

import (
	"fmt"
)

func tips_auto() {
  // identificando os tipos automagicamente
  age := 21
  name := "pauwlo"
  live := true
  
  fmt.Println(age, name, live)
}

func tips_fixed() {
  // aqui as vars são tipadas por mim e seu tipo não pode ser alterado
  // o uso do sufixo "var" indica que essa var pode ser acessada em qualquer parte do package, semelhante ao var vs let em js, o:= é let e var é var
 var age int = 22
 var name string = "pauwlo"
 var live bool = true

  fmt.Println(age, name, live)
}


func main() {
  tips_auto()
  tips_fixed()
}
