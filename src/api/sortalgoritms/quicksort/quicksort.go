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
		sort(array)
	}

	return now, time.Now(), array
}

func sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) - 1

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sort(a[:left])
	sort(a[left+1:])

	return a
}
