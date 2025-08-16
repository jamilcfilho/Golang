package main

import "fmt"

/*
Através do comando IF, descobrir se um numéro é par ou ímpar
*/

func main() {
	for i := 1; i <= 10; i++{
		if i % 2 == 0{
			fmt.Printf("%d = Par \n", i)
		} else {
			fmt.Printf("%d = Ímpar \n", i)
		}
	}
}
