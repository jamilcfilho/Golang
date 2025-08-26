// Select: Funciona como um switch para canais.
// Select permite que guarde operações de vários canais.

// Combinar goroutines e canais com select é um recurso poderoso em Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Recebe os valores "um" e depois "dois" conforme o esperado.
	go func ()  {
		time.Sleep(1 * time.Second)
		c1 <- "um"
	}()
	
	go func ()  {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Recebe:", msg1)

		case msg2 := <-c2:
			fmt.Println("Recebe:", msg2)
		}
	}
}