package pattern

import (
	"sync"
	"time"
)

type Singleton struct {
	Tiempo int64
}

var instancia *Singleton
var once sync.Once

func GetInstancia() *Singleton {
	once.Do(func() {
		instancia = &Singleton{
			time.Now().Unix(),
		}
	})

	return instancia
}
