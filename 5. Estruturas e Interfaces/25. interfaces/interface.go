// Interfaces de forma resumida são uma coleção de métodos

package main

import (
	"fmt"
	"math"
)

// Aqui temos uma interface usada para formas geométricas
type geometria interface{
	area() float64
	perimetro() float64
}

/* Para exemplo de uma interface, será utilizado tipos "Quadrado" e "Circulo", na qual as suas  respectivas áreas e perímetros serão calculadas
	- Quadrado -> área = lado * lado | perímetro = soma dos lados
	- Círculo -> área = (Pi) * raio² | perímetro = 2 * raio * (Pi)
*/

type Quadrado struct{
	lado float64
}

type Circulo struct{
	raio float64
}

func (q Quadrado) area()float64 {
	return q.lado * q.lado
}
func (q Quadrado) perimetro()float64 {
	return 4 * q.lado // Valor 4, vem dos quatro lados do quadrado
}

func (c Circulo) area()float64 {
	return math.Pi * c.raio * c.raio
}
func (c Circulo) perimetro()float64 {
	return 2 * math.Pi * c.raio
}

/* Se a variável tem um tipo interface, então pode-se chamar métodos que estão na interface nomeada.
	Aqui possui uma função geométrica genérica de 'medir' tomando vantagem desse trabalho para qualquer geometria utilizada.*/
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perimetro())
}

func main(){
	q := Quadrado{lado: 10}
	c := Circulo{raio: 5}

	// Os tipos de structs "Quadrado" e "Circulo", ambos implementam a interface 'geometria' então podemos usar instâncias dessas estruturas como argumentos para 'medir'
	medir(q)
	medir(c)
}