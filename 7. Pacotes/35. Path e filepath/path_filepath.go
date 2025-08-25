// Pacote que ajuda a trabalhar com caminhos de arquivos

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	caminho := "pasta/subpasta/arquivo.txt"

	fmt.Println("Diretório:", filepath.Dir(caminho))
	fmt.Println("Arquivo:", filepath.Base(caminho))
	fmt.Println("Extensão:", filepath.Ext(caminho))
}