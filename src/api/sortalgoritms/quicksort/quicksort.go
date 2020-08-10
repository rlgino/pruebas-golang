package quicksort

import (
	"time"

	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms"
)

const (
	quickSort = "Quick Sort"
)

// QuickSort method
type QuickSort struct{}

// New constructor
func New() sortalgoritms.SortMethod {
	return &QuickSort{}
}

// Name of sort method
func (quicksort QuickSort) Name() string {
	return quickSort
}

// Sort method of quick sort
func (quicksort QuickSort) Sort(array []int) (time.Time, time.Time, []int) {
	now := time.Now()

	if len(array) > 0 {
		recursiverSort(array, 0, len(array)-1)
	}

	return now, time.Now(), array
}

func recursiverSort(array []int, start, end int) {
	if (end - start) < 0 {
		return
	}

	pivot := array[end]
	splitIntex := start

	for i := start; i < end; i++ {
		if array[i] < pivot {
			if splitIntex < i {
				aux := array[splitIntex]
				array[splitIntex] = array[i]
				array[i] = aux
			}

			splitIntex++
		}
	}

	array[end] = array[splitIntex]
	array[splitIntex] = pivot

	recursiverSort(array, start, splitIntex-1)
	recursiverSort(array, splitIntex+1, end)
}
