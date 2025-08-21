// Recursão: É a capacidade de uma função chamar ela mesma

package main

import "fmt"

func fatorial(numero int) int {

	if numero == 0 {
		return 1
	}
	return numero * fatorial(numero - 1)
}

func main() {

	fmt.Println(fatorial(5))
}