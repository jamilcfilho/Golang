/* Operador Curto em Go
   Este exemplo demonstra o uso do operador curto de atribuição em Go.
   Na qual utiliza o comando < := > para declarar e inicializar variáveis de forma concisa.
   E o Go já identifica o tipo da variável automaticamente.
   -- Importante: o operador curto só pode ser usado dentro de funções.
   -- Não pode ser usado para declarar variáveis globais.
*/

package main

import "fmt"

// Declarar variáveis globais NÃO pode usar operador curto
const ebulicaoF float64 = 212.0

// Dentro da função é POSSÍVEL usar o operador curto que é o comando :=
func main() {

	tempF := ebulicaoF // Usando operador curto para declarar e inicializar a variável tempF
	tempC := (tempF - 32) * 5 / 9 // Usando operador curto para declarar e inicializar a variável tempC

	fmt.Println("A temperatura de ebulição da água em ºF =",tempF)
	fmt.Println("A temperatura da ebulição da água em ºC =",tempC)

}