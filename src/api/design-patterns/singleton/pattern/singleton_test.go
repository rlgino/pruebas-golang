package pattern_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rlgino/golang-pruebas/src/api/design-patterns/singleton/pattern"
)

func Test_SingletonSameInstance(t *testing.T) {
	mls := pattern.GetInstancia().Tiempo
	time.Sleep(1 * time.Second)

	mls2 := pattern.GetInstancia().Tiempo

	assert.Equal(t, mls, mls2)
}
