package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	fmt.Println("------------------------")

	//bitwise
	/*"Bitwise" é sobre como trabalhar com os bits, que são os dígitos 0 e 1 que formam números em computadores
	Alguns exemplos abaixo de operações bitwise
	*/
	fmt.Println("E =>", a&b)   // 11 & 10= 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
