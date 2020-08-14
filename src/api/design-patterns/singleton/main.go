package main

import (
	"fmt"
	"time"

	"github.com/rlgino/golang-pruebas/src/api/design-patterns/singleton/pattern"
)

func main() {
	fmt.Println("Todas las instancias Singleton tienen que tener el mismo número")

	fmt.Printf("Instancia Singleton: %d\n", pattern.GetInstancia().Tiempo)

	time.Sleep(1 * time.Second)

	fmt.Printf("Instancia Singleton: %d\n", pattern.GetInstancia().Tiempo)

	canalEspera := make(chan int64)
	go func() {
		time.Sleep(1 * time.Second)

		canalEspera <- pattern.GetInstancia().Tiempo
	}()

	fmt.Printf("Instancia Singleton: %d\n", <-canalEspera)
}
