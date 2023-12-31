# Bubble Sort

É um algoritmo de ordenação simples que percorre repetidamente a lista, compara elementos adjacentes e os troca se estiverem na ordem errada. O processo é repetido até que nenhuma troca seja necessária, indicando que a lista está ordenada.

## Passos do Algoritmo

1. Comece do início da lista.
2. Compare o elemento atual com o próximo elemento.
3. Se o elemento atual for maior que o próximo, troque-os.
4. Continue percorrendo a lista até o final.
5. Repita os passos 1-4 até que nenhuma troca seja feita durante uma passagem completa pela lista.

## Utilização

Geralmente utilizado para fins educacionais, devido à sua simplicidade. No entanto, devido à sua ineficiência em termos de desempenho, é raramente utilizado em ambientes de produção para grandes conjuntos de dados.

## Complexidade

- **Complexidade de Tempo (Pior Caso):** O(n^2) - Quadrático.
- **Complexidade de Tempo (Melhor Caso):** O(n) - Linear (quando a lista já está ordenada).
- **Complexidade de Espaço:** O(1) - Constante (não requer espaço adicional, além da variável temporária para trocas).

## Notação Assintótica

A notação assintótica deste algoritmo é O(n^2), indicando que seu desempenho piora quadráticamente com o aumento do tamanho da entrada. Isso ocorre porque, em cada passagem pela lista, são feitas comparações e trocas que dependem do tamanho da lista.


Para testar o algoritimo, inclua o seguinte código no arquivo main.go 

```
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
```

