package main

import "fmt"

// Duas formas de escrever o FOR e obter o mesmo resultado

/*
func main() {
	i:=1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
*/

/*

- Sintaxe básico do FOR

for <inicialização>; <condição>; <pós-execução> {
	<bloco de código>
}

*/

func main() {
	for i := 1 ; i <= 10; i++{
		fmt.Println(i)
	}
}