package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms/bublesort"
	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms/mergesort"
	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms/quicksort"
)

func main() {
	var array [200]int
	size := len(array)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}

	var start, end time.Time
	var elapsed time.Duration

	// Buble sort
	bubleSort := bublesort.New()
	start, end, _ = bubleSort.Sort(array[:])
	elapsed = end.Sub(start)

	fmt.Println(bubleSort.Name())
	fmt.Println("Inicio " + start.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Fin " + end.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Total diferencia: ", elapsed)

	// Merge sort
	mergesort := mergesort.New()
	start, end, _ = mergesort.Sort(array[:])
	elapsed = end.Sub(start)

	fmt.Println(mergesort.Name())
	fmt.Println("Inicio " + start.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Fin " + end.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Total diferencia: ", elapsed)

	// Quick sort
	quicksort := quicksort.New()
	start, end, _ = quicksort.Sort(array[:])
	elapsed = end.Sub(start)

	fmt.Println(quicksort.Name())
	fmt.Println("Inicio " + start.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Fin " + end.Format("2 Jan 2006 15:04:05"))
	fmt.Println("Total diferencia: ", elapsed)
}
