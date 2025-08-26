// Canal: Modo que duas goroutines se comuniquem e sincronizem sua execução.
// Paavra reservada para o channels = <chan>

package main

import (
	"fmt"
)

// Código que demonstra o canal sendo executado e será finalizado somente quando apertar a tecla "Enter"
// func pingar(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "Ping" // Usado para enviar e receber mensagem pelo canal
// 	}
// }

// func imprimir(c chan string) {
// 	for {
// 		msg := <- c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 2)
// 	}
// }

// func main() {
// 	var c chan string = make(chan string)

// 	go pingar(c)
// 	go imprimir(c)

// 	var entrada string
// 	fmt.Scanln(&entrada)
// }


func produtor(out chan<- int) { // só envia
    for i := 0; i < 3; i++ {
        out <- i
    }
    close(out)
}

func consumidor(in <-chan int) { // só recebe
    for v := range in {
        fmt.Println("consumiu:", v)
    }
}

func main() {
    ch := make(chan int)
    go produtor(ch)
    consumidor(ch)
}
