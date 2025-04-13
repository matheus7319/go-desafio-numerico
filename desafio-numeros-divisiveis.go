package main

import "fmt"

func main() {
	startList := 1
	endList := 100 // corrigido aqui

	divisibleForNumber := 3
	fmt.Printf("Iniciando Desafio 1, printar divisiveis por %d :)\n", divisibleForNumber)
	printFromToIfDivisible(startList, endList, divisibleForNumber)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	pinNumber := 3
	panNumber := 5
	fmt.Printf("Iniciando Desafio 2, printar PIN e/ou PAN para divisiveis por %d e %d :)\n", pinNumber, panNumber)
	printPinPanFromTo(startList, endList, pinNumber, panNumber)
}

func printFromToIfDivisible(start int, limit int, divisible int) {
	for i := start; i <= limit; i++ {
		if i%divisible == 0 {
			fmt.Println(i)
		}
	}
}

func printPinPanFromTo(start int, limit int, pinCheck int, panCheck int) {
	for i := start; i <= limit; i++ {
		if i%pinCheck == 0 || i%panCheck == 0 {
			fmt.Printf("Número %d:\n", i)

			if i%pinCheck == 0 {
				fmt.Println("- PIN!")
			}

			if i%panCheck == 0 {
				fmt.Println("- PAN!!")
			}
			fmt.Println("")
		}
	}
}
