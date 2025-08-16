package main

import "fmt"

// While = Enquanto algo for <verdade>, faça determinada ação
func main(){

	i := 0

	for i < 20 {
		fmt.Printf("%d é menor que 20.\n", i)
		i++
	}
}