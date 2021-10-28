package main

import (
  "fmt"
)

/*
 * - Constantes
 * 
 * - São valores imutáveis.
 * - Podem ser tipadas ou não:
 *   - const oi = "Bom dia!"
 *   - const oi string = "Bom dia!"
 * - As não tipadas só teram um tipo atribuido a elas quando forem utilizaadas.
 *   -Ex. qual o tipo de 42? int? uint? float64?
 *   - Ou seja, é uma flexibilidade conveniente. 
 * - Na prática: int, float, string
 *   - const x = y
 *   - const ( x = y )
 * 
 */

const x = 10 // vai continuar sendo indefinido, até ser usado

var (
  y = 10 // vai ser definido como iint no momento da compilaçao
  a int
  b float64
)

func main() {
  fmt.Println("Aula 09, Constantes")
  a = x
  fmt.Printf("%d\t - \t%T\n", a, a) // output: 10 - int
  b = x
  fmt.Printf("%.2f\t - \t%T\n", b, b) // output 10.00 - float64
  // isso funciona pois a constante permanece indefiida até o seu uso, SHOW \o/
}
