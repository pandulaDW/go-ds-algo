package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}

	// assert that bubble sort correctly sort in ascending order
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 9, 10, 11, 12}, BubbleSort(arr, true))

	// assert that bubble sort correctly sort in descending order
	assert.Equal(t, []int{12, 11, 10, 9, 7, 6, 5, 4, 3, 2}, BubbleSort(arr, false))
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}

	for i := 0; i < b.N; i++ {
		BubbleSort(arr, true)
	}
}
