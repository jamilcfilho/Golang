// Exemplo completo utilizando todos as funcionalidades aprendidas

package main

import "fmt"

// Função
func somar(a, b int) int {
	return a + b
}

// Função recursiva
func somaRecursivaAte(limite int) int {
	if limite == 0 {
		return 0
	}
	return limite + somaRecursivaAte(limite - 1)
}

// Função que altera valor com utilização do Ponteiro
func dobrar(numero *int) {
	*numero = *numero * 2
}

func main() {
	// Closure
	acumulador := func () func(int) int {
		total := 0
		return func(valor int) int {
			total += valor
			return total
		}
	}()

	fmt.Println("Acumulador:", acumulador(5)) // 5
	fmt.Println("Acumulador:", acumulador(3)) // 8

	// Recursão
	fmt.Println("Soma recursiva até 5:", somaRecursivaAte(5)) // 15

	// Defer
	defer fmt.Println(">>> Programa finalizado (mensagem do Defer)")

	// Panic e Recover
	func ()  {
		defer func ()  {
			if mensagem := recover(); mensagem != nil {
				fmt.Println("Erro tratado com recover:", mensagem)
			}
		} ()
		fmt.Println("Tentando algo arriscado...")
		panic("Falha critica simulada")
	} ()

	// Ponteiros
	numeroOriginal := 10
	dobrar(&numeroOriginal)
	fmt.Println("Número original dobrado:", numeroOriginal) // 20
}