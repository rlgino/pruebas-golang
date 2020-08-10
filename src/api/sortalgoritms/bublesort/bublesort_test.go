package bublesort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rlgino/golang-pruebas/src/api/sortalgoritms/bublesort"
)

func Test_sort(t *testing.T) {

	arrayTest := []struct {
		name  string
		array []int
		ass   []int
	}{
		{
			name:  "Sort from grather than less",
			array: []int{5, 4, 3, 2, 1},
			ass:   []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Sort from less than grather",
			array: []int{1, 2, 3, 4, 5},
			ass:   []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Sort all equals",
			array: []int{1, 1, 1, 1},
			ass:   []int{1, 1, 1, 1},
		},
		{
			name:  "Empty array",
			array: []int{},
			ass:   []int{},
		},
	}

	bublesort := bublesort.New()
	for _, tt := range arrayTest {
		t.Run(tt.name, func(localTest *testing.T) {
			ass := assert.New(localTest)
			arr := tt.array
			bublesort.Sort(arr)
			ass.EqualValues(arr, tt.ass)
		})
	}
}
