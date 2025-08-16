package main

import "fmt"

const ebulicaoF = 212.0

func main(){

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9
	fmt.Printf("A temperatura da ebulição da água em ºF = %.2f e temperatura de ebulição da água em ºC = %.2f",tempF, tempC)
}