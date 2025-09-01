// Service: Servidor de serviço, que irá prestar algo ao usuário

package main

import (
	"fmt"
	"net/http"
)

/*
Um conceito fundamental em servidores "net/http" são as funções (que estão guardando o nosso pacote http)
*/

/*
As funções que servem como manipuladores, recebem a <http.ResponseWriter> e a <http.Request> como argumentos.
O gravador de resposta é usado para preencher a resposta HTTP. Aqui a nossa resposta é simples, pois é apenas um "Olá mundo!"
*/
func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Olá mundo! \n")
}

/*
Esse próximo manipulador faz algo um pouco mais sofisticado que serve para ler todos os cabeçalhos de solicitação HTTP e ecoa no corpo da resposta
*/
func cabecalho(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalho := range req.Header {
		for _, c := range cabecalho {
			fmt.Fprintf(w, "%v: %v \n", nome, c)
		}
	}
}

/*
Um manipulador (função) é um objeto que implementa <http.Handler>
Uma maneira comum de escrever um manipulador (handler) é utilizando o <http.HandlerFun> adaptando as funções com a assinatura própria.
Registramos nossos manipuladores nas rotas do servidor usando a <http.HandlerFunc> onde ele deverá chamar "/ola" e "/cabecalho".
Ele configura o roteador padrão no pacote "net/http" e recebe uma função como argumento. ("", ola) e ("", cabecalho)
*/
func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalho)
	
	/*
	Finalmente, chamamos o <ListenAndServe> com a porta ":8080" e um manipulador <nil> para que seja usado o roteador padrão que acabamos de configurar.
	*/
	http.ListenAndServe(":8080", nil)
}

/*
Execute o servidor em segundo plano acessando:
1) http://localhost:8080/ola
2) http://localhost:8080/cabecalho
*/