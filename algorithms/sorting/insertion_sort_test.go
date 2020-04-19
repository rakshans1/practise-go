package sorting

import (
	"practise-go/algorithms/sorting/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchMarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

func TestInsertionSortString(t *testing.T) {
	sorttest.TestString(t, InsertionSortString)
}
