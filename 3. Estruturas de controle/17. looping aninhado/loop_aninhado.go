package main

import "fmt"

// É possível ter loops dentro de loops, por exemplo, para trabalhar com matrizes.
func main(){

	for i := 1; i < 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d \n", i, j)
		}
	}
}