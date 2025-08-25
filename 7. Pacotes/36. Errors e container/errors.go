// Pacote que permite tratar/manipular erros

package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Divisão por zero não é permitida")
	}
	return dividendo / divisor, nil
}

func main() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)
}