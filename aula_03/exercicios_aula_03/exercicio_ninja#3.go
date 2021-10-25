package main

import (
  "fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true


func main()  {

  fmt.Println("exercicio 3, package level, tipos and Sprintf")
  fmt.Println("==============================================")
  s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
  fmt.Println(s)
}

