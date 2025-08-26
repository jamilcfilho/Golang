// Goroutines: Funções capaz de executar outras funções que rodam concorrentemente
// Palavra resrevada para a goroutine é <go>

package main

import (
	"fmt"
	"time"
)

func mensagem(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Executando a função normalmente
	mensagem("Executando a função principal")

	// Executando como goroutine
	go mensagem("Executando a goroutine")

	// Esperar um pouco para a goroutine terminar
	time.Sleep(2 * time.Second)
}