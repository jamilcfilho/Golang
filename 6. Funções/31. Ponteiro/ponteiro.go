/*

Ponteiro: Em Go armazena um valor na memória, mas não é o valor propriamente escrito, é uma cópia. Por isso se quisermos modificar esse valor, utilizamos um "ponteiro <variávelPtr>"".
& -> endereço da variável
* -> valor armazenado no endereço

*/

package main

import "fmt"

func dobrar(xPtr *int) {
	*xPtr = *xPtr * 2
}

func main() {
	numero := 10
	dobrar(&numero)
	fmt.Println(numero)
}