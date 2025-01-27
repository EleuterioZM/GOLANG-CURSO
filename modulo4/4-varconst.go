package main

import (
	"fmt"
)

/*
	func main() {
		variaveis()
	}
*/
func variaveis() {
	// variaveis
	//var + nome da variavel + Tipo
	var nome string
	nome = "Eleuterio"
	fmt.Println(nome)

	nome = "Raul"
	fmt.Println(nome)

	var idade, ano int = 15, 2025
	
	fmt.Println(idade)
	fmt.Println(ano)
	
	// Outra forma, inferencia de tipos
ghost := 007
ghost  = 5
fmt.Print(ghost)

//constantes

const idade_eleuetrio = 4

fmt.Print(idade_eleuetrio)
}
