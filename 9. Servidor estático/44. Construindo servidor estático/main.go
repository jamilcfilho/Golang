/*
Será criado um site estático em exemplo simples mas real, na qual o código fará o seguinte processo:
- Servir arquivos HTML de um local específico no disco
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Cria um file server apontando para a pasta "estatico"
	fs := http.FileServer(http.Dir("./estatico"))
	
	// Roteia todas as requisições para o file server
	http.Handle("/", fs)

	fmt.Println("Servidor estático rodando em http://localhost:3000")
	log.Print("Listening on: 3000...")
	
	// Inicia o servidor e colocando o destino de porta <3000> (para esse exemplo) e indicando um err para caso ocorra um erro na criação
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor", err)
	}
}