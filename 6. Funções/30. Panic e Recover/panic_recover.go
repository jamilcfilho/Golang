// Panic: Aparece quando possui um erros graves de programação ou erro de execução pelo tempo longo
// Recover: Serve para se recuperar de um <panic>

package main

import "fmt"

func arriscar() {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Recuperação de:", r)
		}
	} ()
	fmt.Println("Antes do panic")
	
	panic("Algo deu errado")
	
	fmt.Println("Depois do panic") // Não será executado
}

func main() {
	arriscar()
	fmt.Println("Programa continua...")
}