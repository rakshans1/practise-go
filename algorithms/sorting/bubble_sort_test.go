package sorting

import (
	"practise-go/algorithms/sorting/sorttest"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}

func BenchMarkBubbleSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, BubbleSortInt)
}

func TestBubbleSortString(t *testing.T) {
	sorttest.TestString(t, BubbleSortString)
}
