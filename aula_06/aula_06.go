package main

import (
  "fmt"
  "runtime"
)

/*
 * - int vs float: Números inteiros vs. números com frações
 * - int & uint => "implementation-specific sizes"
 * - Todos os tipos numericos são distintos, exceto:
 *   - byte = uint8 (2^8)
 *   - rune = int32 (UTF-8) (2^32)
 * - Tipos são únicos
 *   - Go é uma linguagem estática
 *   - int e int32 não são a mesma coisa
 *   - para "misturá-los" é necessário conversão
 * - Regra geral: Use sempre int
 * 
 * - Floting Point:
 *   - Números racionais ou reais
 *   - Regra gereall: Use somente float64
 * -Na prática:
 *   - Defaults com :=
 *   - Tipagem com var
 */

  func main() {
    fmt.Println("aula 06, Tipos númericos e Package runtime")
    fmt.Println(runtime.GOOS) // Mostra o sistema operacional
    fmt.Println(runtime.GOARCH) // Processador 
}

