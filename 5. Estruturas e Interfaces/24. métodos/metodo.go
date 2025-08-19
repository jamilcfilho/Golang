/* Método é uma função associada a um tipo particular, geralmente a uma struct.
Isto é, em Programação Orientada a Objeto, objeto é um valor(variável) e o método é atribuir uma função associada a esse objeto */

package main

import "fmt"

type retangulo struct{
	base, altura int
}

// Esse método "area" possui um tipo "retangulo" em formato de ponteiro com o síbolo <*>
func (r *retangulo) area()int{
	return r.base * r.altura
}

/* Métodos poder ser definidos por qualquer tipo de receptor ponteiro ou valor.
	Aqui temos um exemplo de receptor de valor. */
func (r retangulo) perimetro() int{
	return 2 * r.base + 2 * r.altura
}

func main(){
	r := retangulo{base: 10, altura: 5}
	p := Pessoa{Nome: "Ana", Idade: 30}

	
	// Aqui chama-se os 2 métodos definidor para a nossa struct
	fmt.Println("Área: ", r.area())
	fmt.Println("Perímetro: ", r.perimetro())
	
	// Aqui chama-se os métodos do que foi criado posteriormente
	p.Saudacao()
	p.MostrarIdade()
	p.FazerAniversario()
}


// Outro exemplo de método para melhor fixação do conteúdo aprendido
type Pessoa struct{
	Nome string
	Idade int
}

func (p *Pessoa) Saudacao(){
	fmt.Println("Olá,", p.Nome)
}

func (p *Pessoa) MostrarIdade(){
	fmt.Println("Idade:", p.Idade)
}

// Possibilidade de incrementar o valor do elemento = Isso é possível porque o método esta como um ponteiro <*>, no caso o <*Pessoa>
func (p *Pessoa) FazerAniversario(){
	p.Idade++
	fmt.Printf("Hoje é seu aniversário, e você esta fazendo %d anos\n\n", p.Idade)
}
// Para utilização desses novos métodos eles, precisam ser incluídos dentro da função principal