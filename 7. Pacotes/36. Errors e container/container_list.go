// Pacote que implementa uma lista duplamente encadeada, útil quando há muitas inserções/remoções

package main

import (
	"container/list"
	"fmt"
)

func main() {
	linguagens := list.New()

	linguagens.PushBack("Go")
	linguagens.PushBack("SQL")
	linguagens.PushFront("Python")

	for item := linguagens.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}