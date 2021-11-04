package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula 14, loops: nested loop (reptição hierárquica)")
	for i := 01; i <= 12; i++ {
		fmt.Printf("\nMês: %v\n", i)

		for j := 01; j <= 30; j++ {
			fmt.Printf("dia: %v ", j)
		}

		fmt.Println("")
	}
}
