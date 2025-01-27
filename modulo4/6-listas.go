package main

import "fmt"

func main() {
	//Arrays
	var arr = [4]string{"Go", "is", "awesome"}
	arr[3] = "THE_GHOST"
	fmt.Println(arr) // Saída: [Go is awesome]


	numPrimos := [4]int{2, 3, 5, 7}
	fmt.Println(len(numPrimos))
	fmt.Println(numPrimos[0:3])
	fmt.Println(numPrimos[1:])

	//Slices
	slice :=make([]string, 5) 
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	slice = append(slice, "oi")
	fmt.Println(slice)
	
}

//Arrays

//Definição: Um array é uma coleção fixa e homogênea de elementos. O tamanho do array é definido na sua declaração e não pode ser alterado.
//var arr [5]int // Array de inteiros com tamanho 5
//arr[0] = 10    // Atribuição de valor ao primeiro índice
//fmt.Println(arr) // Saída: [10 0 0 0 0]

//----------------------------------------------------------------------

//Slices

//Definição: Um slice é uma abstração de um array. Ele é dinâmico e flexível, permitindo redimensionamento durante a execução.

//var slice []int          // Slice de inteiros
//slice = append(slice, 1) // Adiciona 1 ao slice
//fmt.Println(slice)       // Saída: [1]
