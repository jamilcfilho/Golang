// Consumindo uma API píblica, na qual é a PokeAPI para buscar informações sobre Pokemons da região de Kanto.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PokemonSpecies representa os dados básicos do Pokemon
type PokemonSpecies struct {
	Nome string `json:"name"`
	URL string `json:"url"`
}

// PokemonEntry representa cada entrada de Pokemon na lista
type PokemonEntry struct {
	EntryNumber int `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

// PokeApiResponse é a estrutura principal que recebe a resposta completa
type PokeApiResponse struct {
	PokemonEntry []PokemonEntry `json:"pokemon_entries"`
}

func main() {
	// Fazer a requisição GET para a URL da API
	resp, err := http.Get("https://pokeapi.co/api/v2/pokedex/Kanto/")
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}

	// Garante que o corpo da resposta será fechado ao final da aplicação = *Boa prática*
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	//Imprime todo o corpo da resposta do arquivo JSON
	//fmt.Println(string(body))

	// Desempacotamento da API que está no formato JSON
	var pokedex PokeApiResponse //Cria uma variável do tipo da nossa struct principal

	err = json.Unmarshal(body, &pokedex)
	if err != nil {
		fmt.Println("Erro ao desempacotar o JSON:", err)
		return
	}

	//Listando os dados
	fmt.Println("\n--- Os 10 primeiros Pokemons da Pokedéx de Kanto ---")
	for i, pokemon := range pokedex.PokemonEntry {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s\n", pokemon.EntryNumber, pokemon.PokemonSpecies.Nome)
	}
}