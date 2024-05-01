package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reverse retorna a string invertida.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Digite a string:")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	// Remover espaços em branco extras e nova linha
	input = strings.TrimSpace(input)

	// Verificar se a entrada não está vazia
	if input == "" {
		fmt.Println("A entrada está vazia.")
		return
	}

	reversedStr := Reverse(input)
	fmt.Println("String invertida:", reversedStr)

	if strings.EqualFold(strings.ReplaceAll(input, " ", ""), Reverse(strings.ReplaceAll(input, " ", ""))) {
		fmt.Println("É um palíndromo.")
	} else {
		fmt.Println("Não é um palíndromo.")
	}
}
