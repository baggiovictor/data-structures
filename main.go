package main

import (
	"data-structures/structures/sorting/bubblesort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)

	fmt.Println("Array antes da ordenação:", arr)

	executions, duration := bubblesort.BubbleSort(arr)

	// Imprimir o array após a ordenação
	fmt.Println("Array após a ordenação:", arr)

	fmt.Printf("Tempo de execução: %v\n", duration)
	fmt.Printf("Número total de execuções: %d\n", executions)
}
