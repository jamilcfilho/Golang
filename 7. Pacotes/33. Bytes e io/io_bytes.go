// Pacote que lida com entrada e saída de dados
// io = fornece interfaces para leitura e escrita
// bytes = permite manipular dados binários em memória

package main

import (
	"bytes"
	"fmt"
	"io"
)

// Código para simular leitura de um arquivo em pequenos blocos
func main() {
	mensagem := "Estudando Go passo a passo"
	leitor := bytes.NewReader([]byte(mensagem))

	buffer := make([]byte, 10) // Leitura em blocos de 10 bytes (lê 10 "letras" e imprime e assim até terminar a frase)
	for {
		n, err := leitor.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}