// Arrays são uma sequência numerada de elementos, do mesmo tipo e tem um tamanho determinado.

package main

import "fmt"

func main (){

	var notas = [5]float64 {5.3, 8, 4.2, 2.1, 7.8}

	var media float64 = 0
	for i := 0; i < len(notas); i++ { // len() para poder "contar" os elementos e facilitar o código
		media += notas[i]
	}

	fmt.Println(media / float64(len(notas))) // Precisa utilizar o comando de "float64" junto do len() para transformar os tipos e manter eles iguais, se não o programa da erro.
}