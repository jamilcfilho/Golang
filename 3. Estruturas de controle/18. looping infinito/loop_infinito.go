package main

import "fmt"

// Looping infinito serve para servidores, HTTP e para processos que precisam ser contínuos.
func main (){

	for {
		fmt.Println("Executando...")
		break // Sem o comando BREAK, o sistema ficaria executando de forma infinita (isso é para esse sistema não se sobrecarregar, porém DEVE SER RETIRADO PARA EXECUTAR DE FORMA INFINITA)
	}
}