// Pacote que permite interação com o sistema operacional: Arquivos, dieretórios, variáveis de ambiente.

package main

import (
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Create("meuarquivo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer arquivo.Close()

	arquivo.WriteString("Go é muito prático! \n")
	fmt.Println("Arquivos criado com sucesso!")
}