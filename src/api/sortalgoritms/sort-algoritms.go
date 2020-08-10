package sortalgoritms

import (
	"time"
)

// SortMethod interface to implement into different algorinthms
type SortMethod interface {
	Sort(array []int) (time.Time, time.Time, []int)
	Name() string
}
