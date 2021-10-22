package main

import "fmt"


func world() {
  fmt.Println("- it's me, world!")
} 


func scene() {
  fmt.Println("And, scene")
}

func main() {
  fmt.Println("- Hello?")
  fmt.Println("- Hi.")
  fmt.Println("- Who is there?")
  fmt.Println("...")
  world()
  fmt.Println("- Ah, hello, world.")
  scene()
}
