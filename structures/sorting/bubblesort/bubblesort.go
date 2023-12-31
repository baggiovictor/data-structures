package bubblesort

import "time"

func BubbleSort(arr []int) (int, time.Duration) {
	n := len(arr)
	executions := 0

	startTime := time.Now()

	// Percorre o array n vezes (inicio do array até o penúltimo elemento)
	for i := 0; i < n; i++ {
		// Percorre o array da primeira posição até a posição n - i - 1
		for j := 0; j < n-i-1; j++ {
			executions++
			// Se o elemento atual for maior que o próximo, troca os elementos de posição
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // swap
			}
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	return executions, duration
}
