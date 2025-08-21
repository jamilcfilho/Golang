// Função, também pode ser chamada de procedimento ou sub rotina (em livros/artigos, etc.)
// Parte do código que recebe uma entrada (parâmetros) e retorna valores.

package main

import "fmt"

// Programa que calcula a média de notas de uma sala de aula
func media (lista []float64)float64 { // Cria-se essa função que vai "conversar" com a função <main>. Pode ser feito várias funções isoladas para que nas boas práticas o código fique fácil de dar manutenção.

	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() {

	lista := []float64 {98, 93, 77, 82, 83} // Lista de notas dos alunos

	fmt.Println(media(lista))
}