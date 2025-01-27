package main

import "fmt"

func main() {
	soma_valores := soma(42, 12)
	fmt.Println(soma_valores)

	Conc_nome := nomes("Eleuterio",  "_Zacarias")
	fmt.Println(Conc_nome)
	

	sub :=subtraccao(12, 2)
	fmt.Println(sub)

}
func soma(x int, y int) int {
	return x + y
}

func subtraccao(x int, y int) int {
	return x - y
}

func nomes(nome_1 string, nome_2 string) string {
	return nome_1 + nome_2
}
