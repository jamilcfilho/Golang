package main

import "fmt"

func main() {
	fmt.Println(len("Hello World!")) // Define o tamanho da string (conta o "espaço" também)
	fmt.Println("Hello World!"[2]) // Acessa o terceiro caractere da string (começa do zero) - Porém o Go considera o caractere como um byte por isso a letra "l" saiu com valor 108
	fmt.Println("Hello" + "world!") // Concatena duas strings
}