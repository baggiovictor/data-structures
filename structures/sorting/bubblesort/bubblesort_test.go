// bubblesort_test.go
package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Cenário 1: Array desordenado
	arr1 := []int{5, 3, 1, 4, 2}
	expected1 := []int{1, 2, 3, 4, 5}
	testBubbleSort(t, arr1, expected1)

	// Cenário 2: Array já ordenado
	arr2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5} // Não deve haver alteração
	testBubbleSort(t, arr2, expected2)

	// Cenário 3: Array com elementos repetidos
	arr3 := []int{3, 2, 1, 3, 1, 2}
	expected3 := []int{1, 1, 2, 2, 3, 3}
	testBubbleSort(t, arr3, expected3)
}

func testBubbleSort(t *testing.T, input, expected []int) {
	executions, duration := BubbleSort(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Erro no teste. Array esperado: %v, Array obtido: %v", expected, input)
	}

	t.Logf("Execuções: %d, Tempo de execução: %v", executions, duration)
}
