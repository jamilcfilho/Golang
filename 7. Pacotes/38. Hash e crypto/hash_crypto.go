// Pacote que permite aplicar funções de hash e algoritmos de criptografia (é muito utilizado em segurança de dados)

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	senha := "minhaSenhaSegura123"
	fmt.Println("Sem criptografia: \n", senha)

	hash := sha256.Sum256([]byte(senha))
	fmt.Println("Com criptografia SHA-256: \n", hash)
}