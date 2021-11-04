package main

import (
  "fmt"
)

func fizBuzz(num int) {

  for i := 1; i <= num; i++ {
    switch {

      case i%15 == 0:
        fmt.Printf("%v:\t FizzBuzz\n", i)

      case i%3 == 0:
        fmt.Printf("%v:\t Fizz\n", i)

      case i%5 == 0:
        fmt.Printf("%v:\t Buzz\n", i)

      default:
        fmt.Printf("%v:\t :/\n", i)
    }
  }
}

func main() {
  fmt.Println("FizzBuzz, i%3 == 0 Fizz, i%5 == 0 Buzz, i%3 && i%5 == 0 FizzBuzz.")
  fizBuzz(200)
}
