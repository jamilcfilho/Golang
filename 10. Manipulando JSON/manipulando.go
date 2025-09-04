// Criando um arquivo JSON a partir de uma struct e salvando no formato .json e realizando a manipulação do mesmo através de abertura e leitura do arquivo criado

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Struct para criar um usuário
type Usuario struct {
	Nome string `json:"Nome"` // Lembrar: Deve-se utilizar o sinal de crase<`> dentro da struct
	Tipo string `json:"Tipo"`
	Idade int `json:"Idade"`
	Ativo bool `json:"Ativo"`
	Social Social `json:"Social"` // Recebendo as struct Social
}

// Struct que armazena as redes sociais para colocar na struct de Usuario a cima
type Social struct {
	Facebook string `json:"Facebook"`
	Instagran string `json:"Instagran"`
}

func main() {
	// Criando uma slice de Usuario, para conseguir armazenar mais usuários
	usuario := []Usuario{
		{Nome: "Maria", Tipo: "Leitora", Idade: 30, Ativo: true, Social: Social{Facebook: "facebook/Maria", Instagran: "instagran/Maria"}},
		{Nome: "Joana", Tipo: "Autora", Idade: 22, Ativo: true, Social: Social{Facebook: "facebook/Joana", Instagran: "instagran/Joana"}},
		{Nome: "Pedro", Tipo: "Leitor", Idade: 36, Ativo: false},
	}
	// Converte a struct para JSON
	dados, err := json.MarshalIndent(usuario,"", " ")
	if err != nil {
		fmt.Println("Erro ao gerar JSON:", err)
		return
	}

	// Cria e gera o arquivo JSON
	err = os.WriteFile("usuario.json", dados, 0644)
	if err != nil {
		fmt.Println("Erro salvar o arquivo", err)
		return
	}

	fmt.Println("Arquivo usuario.json criado com sucesso!")

	// Realiza a abertura do aquivo JSON criado
	jsonFile, err := os.Open("usuario.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}

	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()

	// Realiza a leitura do arquivo criado
	dados, err = os.ReadFile("usuario.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo", err)
		return
	}

	// Converte JSON para struct
	err = json.Unmarshal(dados, &usuario)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	// Realiza a leitura de todos os usuários de forma direta porém com informações mais "bagunçadas" pois ficam unidas sem separação
	fmt.Println("Usuários carregados:", usuario)

	// Imprime de um jeito mais configurado e organizado
	for i := 0; i < len(usuario); i++ {
	fmt.Println("Nome:" + usuario[i].Nome)
	fmt.Println("Tipo:" + usuario[i].Tipo)
	fmt.Println("Idade:" + strconv.Itoa(usuario[i].Idade))
	fmt.Println("Ativo:" + strconv.FormatBool(usuario[i].Ativo))
	fmt.Println("Redes Sociais: Facebook - " + usuario[i].Social.Facebook + "\nInstagran - " + usuario[i].Social.Instagran)
	fmt.Println("-------")
	}

}