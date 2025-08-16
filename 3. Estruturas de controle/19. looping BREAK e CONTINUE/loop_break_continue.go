package main

import "fmt"

// Interromper ou pular iterações, deve ser realizado utilizando os comandos BREAK e CONTINUE
func main (){

	for i := 1; i < 5; i++ {
		if i == 3 {
			continue // Pula o número 3 e continua a execução
		}
		if i == 5 {
			break // Para o loop imediatamente, no caso no momento que atinge o valor de 5 e não executa a função
		}
		fmt.Println(i)
	}
}