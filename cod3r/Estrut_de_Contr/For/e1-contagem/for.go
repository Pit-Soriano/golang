package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 0; i <= 10; i += 1 {
		fmt.Printf("\n%d", i)
	}
}
