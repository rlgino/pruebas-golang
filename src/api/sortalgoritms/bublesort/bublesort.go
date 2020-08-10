package bublesort

import (
	"time"

	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms"
)

const (
	bubleSort = "Buble Sort"
)

// Bublesort method
type Bublesort struct{}

// New constructor
func New() sortalgoritms.SortMethod {
	return &Bublesort{}
}

// Name of sort method
func (bublesort Bublesort) Name() string {
	return bubleSort
}

// Sort method of buble sort
func (bublesort Bublesort) Sort(array []int) (time.Time, time.Time, []int) {
	now := time.Now()
	for x := 0; x < len(array); x++ {
		for i := 0; i < len(array); i++ {
			if array[i] > array[x] {
				aux := array[i]
				array[i] = array[x]
				array[x] = aux
			}
		}
	}

	return now, time.Now(), array
}
