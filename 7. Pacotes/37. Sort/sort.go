// Pacote que organiza os slices(fatias)

package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := []int {5, 1, 9, 3}
	sort.Ints(numeros)
	fmt.Println("Ordenados:", numeros)

	nomes := []string {"Carlos", "Ana", "Beatriz"}
	sort.Strings(nomes)
	fmt.Println("Ordenados:", nomes)
}