package insertionsort

import "time"

func InsertionSort(arr []int) (int, time.Duration) {
	n := len(arr)
	executions := 0

	startTime := time.Now()

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			executions++
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	return executions, duration
}
