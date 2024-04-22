package main

import (
	"fmt"
)

func Reverse(primeiraString string) string {
	runes := []rune(primeiraString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var primeiraString string
	fmt.Println("Digite a primeia string")
	fmt.Scanln(&primeiraString)

	Reversa_Str := Reverse(primeiraString)
	fmt.Println(Reversa_Str)

	if primeiraString == Reversa_Str {
		fmt.Println("é um palindromo")
	} else {
		fmt.Println("não é um palindromo")
	}
}
