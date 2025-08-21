/* Defer: Escalona as funções, no caso ela "agenda" a execução de uma função *para o final da função atual*.
Muito utilizado para liberar recursos (arquivos, conexões, etc.)
Os <defer> são executados em ordem inversa (como em uma pilha)
*/

package main

import "fmt"

func inicio() {
	fmt.Println("Iníciando o programa...")
}

func defer1() {
	fmt.Println("Defer 1")
}

func defer2() {
	fmt.Println("Defer 2")
}

func fim() {
	fmt.Println("Finalizando o programa...")
}

func main() {
	inicio()
	defer defer1()
	defer defer2()
	fim()

}