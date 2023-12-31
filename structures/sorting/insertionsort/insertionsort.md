# Insertion Sort

## Descrição

É um algoritmo de ordenação simples que constrói uma sequência ordenada de elementos, um de cada vez. Ele é muito eficiente para listas pequenas e para listas que estão quase ordenadas inicialmente.

## Passos do Algoritmo

### Partição Inicial

- O primeiro elemento é considerado uma lista ordenada de um único elemento.

### Inserção

- Para cada elemento subsequente, é inserido na posição correta na lista ordenada até o momento.
- O elemento é comparado com os elementos na lista ordenada, e é deslocado para a posição correta, empurrando os elementos maiores uma posição para a direita.

### Repetição

- O processo é repetido até que todos os elementos estejam na posição correta.

## Utilização

O Insertion Sort é eficiente para listas pequenas ou quase ordenadas. No entanto, para grandes conjuntos de dados, algoritmos de ordenação mais eficientes como QuickSort ou MergeSort são geralmente preferidos.

## Complexidade

- **Complexidade de Tempo (Pior Caso):** \(O(n^2)\) - Quadrático.
- **Complexidade de Tempo (Melhor Caso):** \(O(n)\) - Linear (quando a lista já está quase ordenada).
- **Complexidade de Espaço:** \(O(1)\) - Constante.

## Notação Assintótica

A notação assintótica do Insertion Sort é \(O(n^2)\), indicando que seu desempenho piora quadráticamente com o aumento do tamanho da entrada.

```
import (
	"data-structures/structures/sorting/insertionsort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10000)

	fmt.Println("Array antes da ordenação:", arr)

	executions, duration := insertionsort.InsertionSort(arr)

	// Imprimir o array após a ordenação
	fmt.Println("Array após a ordenação:", arr)

	fmt.Printf("Tempo de execução: %v\n", duration)
	fmt.Printf("Número total de execuções: %d\n", executions)
}
```