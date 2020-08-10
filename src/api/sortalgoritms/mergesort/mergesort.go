package mergesort

import (
	"time"

	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms"
)

const (
	mergeSortName = "Merge Sort"
)

// Mergesort method
type Mergesort struct{}

// New constructor
func New() sortalgoritms.SortMethod {
	return &Mergesort{}
}

// Name of the method
func (Mergesort) Name() string {
	return mergeSortName
}

//Sort method to execute mergesort method
func (mergesort Mergesort) Sort(array []int) (time.Time, time.Time, []int) {
	start := time.Now()

	res := mergeSort(array)

	return start, time.Now(), res
}

func mergeSort(array []int) []int {
	size := len(array)

	if size == 1 {
		return array
	}

	middle := size / 2
	var (
		leftArray  = array[0:middle]
		rigthArray = array[middle:size]
	)

	return merge(mergeSort(leftArray), mergeSort(rigthArray))
}

func merge(left []int, rigth []int) (result []int) {
	result = make([]int, len(left)+len(rigth))

	i := 0
	for len(rigth) > 0 && len(left) > 0 {
		if rigth[0] > left[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = rigth[0]
			rigth = rigth[1:]
		}
		i++
	}

	for j := 0; j < len(rigth); j++ {
		result[i] = rigth[j]
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	return
}
