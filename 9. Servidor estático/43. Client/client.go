/*
Client: É um servidor que esta relacionado ao que o usuário precisa.
Nesse código vamos emitir uma solicitação de um cliente a um servidor HTTP
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*
Emita uma solicitação HTTP GET para um servidor.
<http.Get> é um atalho conveniente para criar um client server objeto e chamar seu (Get) método.
*/
func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Ele usa o objeto que tem configurações padrão úteis, como predefinição do cliente
	fmt.Println("Response status:", resp.Status) 
	/* Imprime o status da resposta HTTP. 
	O HTTP é o protocolo responsável por fazer a comunicação entre o cliente e o servidor. 
	Dessa forma, a cada "solicitação" feita, o HTTP responde se você obteve sucesso ou não, se há algum erro na página, etc. Estas mensagens de erro são os "status HTTP".
	Por exemplo, os erros de HTTP mais comuns são: erro http 404; erro http 500; erro http 403
	*/

	// Imprimir as primeiras 5 linhas do corpo da resposta, utilizando a biblioteca <bufio>
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}