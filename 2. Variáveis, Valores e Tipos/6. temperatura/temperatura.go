// Declarar pacote principal
package main

// Importar função <fmt>
import "fmt"

// Declaração de variáveis globais
// Declaração da variável CONST do ponto de ebulição da água em Fahrenheit
const ebulicaoF = 212.0


// Função principal <main>
func main() {

	var tempF = ebulicaoF // Variável do valor da temperatura em Fahrenheit
	var tempC = (tempF - 32) * 5 / 9 // Conversão de Fahrenheit para Celsius

	// Apareça na tela o resultado da conversão
	fmt.Println("A temperatura de ebulição da água em ºF = ", tempF)
	fmt.Println("A temperatura de ebuklição da água em ºC = ",tempC)

}